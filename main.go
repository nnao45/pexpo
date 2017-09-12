package main

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/dariubs/percent"
	"github.com/nsf/termbox-go"
	"github.com/tatsushid/go-fastping"
	"github.com/tlorens/go-ibgetkey"
	//	"io/ioutil"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/*This global counter*/
var i int // "i" is all pings count.
var j int // "j" is all pings "per host" count.

/*This global buffer*/
var pbf bytes.Buffer // pbf is ping-list(textfile -> buffer).
var hbf bytes.Buffer // hbf is ping loss mapping to host.

const layout = "2006 Jan 02 15:04:05.000Z07:00 JST"

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func addog(text string, filename string) {
	var writer *bufio.Writer
	text_data := []byte(text)

	write_file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	writer = bufio.NewWriter(write_file)
	writer.Write(text_data)
	writer.Flush()
	fatal(err)
	defer write_file.Close()
}

func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func drawBlue(x, y int, str string) {
	termbox.SetOutputMode(termbox.Output256)
	color := termbox.Attribute(21 + 1)
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func drawRed(x, y int, str string) {
	termbox.SetOutputMode(termbox.Output256)
	color := termbox.Attribute(196 + 1)
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func drawGreen(x, y int, str string) {
	termbox.SetOutputMode(termbox.Output256)
	color := termbox.Attribute(48 + 1)
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func Pinger(host string, index int) (s string, i int) {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", host)
	if err != nil {
		termbox.Close()
		panic(err)
	}
	p.AddIPAddr(ra)

	var out string
	var res string
	receiver := make(chan string, 10000)
	go func() {
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			//out = "Host: " + host + " IP Addr: " + addr.String() + " receive, RTT: " + rtt.String() + "\n"
			out = "Host: " + host + " receive, RTT: " + rtt.String() + "\n"
			receiver <- out
		}
	}()
	p.OnIdle = func() {
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case res = <-receiver:
			return res, 0
		case <-time.After(time.Second):
			res = "Host: " + host + " ping faild...\n"
			fres := strconv.Itoa(index) + "\n"
			hbf.WriteString(fres)
			return res, 1
		}
	}
}

func draw() {
	drawHostList()
	drawLine(0, 0, "Press q to exit.")
	index := 2
	pscanner := bufio.NewScanner(strings.NewReader(pbf.String()))
	for pscanner.Scan() {
		ps := pscanner.Text()
		res, flag := Pinger(ps, index)
		if flag == 0 {
			drawBlue(2, i+2, fmt.Sprintf("%v", "o"))
		} else if flag == 1 {
			drawRed(2, i+2, fmt.Sprintf("%v", "x"))
		}
		drawLine(4, i+2, fmt.Sprintf("%v", res))
		drawGreen(80, index, fmt.Sprintf("%.2f", Round(percent.PercentOf(drawLoss(index), j), 2)))
		drawGreen(86, index, fmt.Sprintf("(%v loss)", drawLoss(index)))
		t := time.Now()
		drawLine(2, 1, fmt.Sprintf("date: %v", t.Format(layout)))
		termbox.Flush()
		i++
		index++
		if err := pscanner.Err(); err != nil {
			panic(err)
		}
	}
}

func drawHostList() {
	hi := 2
	drawLine(60, hi-1, fmt.Sprintf("%v", "HOST"))
	drawLine(80, hi-1, fmt.Sprintf("%v", "LOSS"))
	scanner := bufio.NewScanner(strings.NewReader(pbf.String()))
	for scanner.Scan() {
		s := scanner.Text()
		drawGreen(60, hi, fmt.Sprintf("%v", s))
		if j <= 1 {
			drawGreen(80, hi, fmt.Sprintf("%v", "0.000"))
			drawGreen(86, hi, fmt.Sprintf("%v", "(0 loss)"))
		}
		hi++
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

}

func drawLoss(index int) int {
	var c int
	scanner := bufio.NewScanner(strings.NewReader(hbf.String()))
	for scanner.Scan() {
		s := scanner.Text()
		if s == strconv.Itoa(index) {
			c++
		}
	}
	return c
}

func pollEvent() {
	kill := make(chan bool)
	finished := make(chan bool)
	go killPing(kill, finished)
	targetkey := "q"
	t := int(targetkey[0])
loop:
	for {
		input := keyboard.ReadKey()
		select {
		case <-finished:
			break loop
		default:
			if input == t {
				kill <- true
				break loop
			}
		}
	}
}

func killPing(kill, finished chan bool) {
	for {
		select {
		case <-kill:
			finished <- true
			return
		default:
			j++
			draw()
		}

	}
	/*
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					return
				}
			}
		}*/
}

func init() {
	pl, err := os.Open("ping-list")
	fatal(err)
	defer pl.Close()
	scanner := bufio.NewScanner(pl)
	for scanner.Scan() {
		s := scanner.Text() + "\n"
		pbf.WriteString(s)
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	pollEvent()
}
