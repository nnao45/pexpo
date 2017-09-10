package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tlorens/go-ibgetkey"
	"github.com/nsf/termbox-go"
	"github.com/tatsushid/go-fastping"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var i int
var h int
var rbf bytes.Buffer
var pbf bytes.Buffer

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func bfdog(text string) string {
	var b bytes.Buffer
	b.WriteString(text)
	return b.String()
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func ifExists(filename string) {
	if !exists(filename) {
		f, err := os.Create(filename)
		fatal(err)
		defer f.Close()
	}
	return
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

func cat(filename string) string {
	buff, err := ioutil.ReadFile(filename)
	fatal(err)
	return string(buff)
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

func Pinger(host string) string {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", host)
	fatal(err)
	p.AddIPAddr(ra)
	var out string
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		out = "Host: " + host + " IP Addr: " + addr.String() + " receive, RTT: " + rtt.String() + "\n"
	}
	p.OnIdle = func() {}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
	rbf.WriteString(out)
	return out

}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawLine(0, 0, "Press q to exit.")
	if i >= 2 {
		rscanner := bufio.NewScanner(strings.NewReader(rbf.String()))
		for rscanner.Scan() {
			rs := rscanner.Text()
			drawBlue(2, h, fmt.Sprintf("%v", "o"))
			drawLine(4, h, fmt.Sprintf("%v", rs))
			h++
			if err := rscanner.Err(); err != nil {
				panic(err)
			}
			i = h
		}
		pscanner := bufio.NewScanner(strings.NewReader(pbf.String()))
		for pscanner.Scan() {
			ps := pscanner.Text()
			drawBlue(2, i, fmt.Sprintf("%v", "o"))
			drawLine(4, i, fmt.Sprintf("%v", Pinger(ps)))
			drawLine(2, 1, fmt.Sprintf("date: %v", time.Now()))
			termbox.Flush()
			i++
			if err := pscanner.Err(); err != nil {
				panic(err)
			}
		}

	}
	termbox.Flush()
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



func killPing(kill, finished chan bool){
	for{
		select {
		case <-kill:
			finished <- true
			return
		default:
			i = 2
			h = 2
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
