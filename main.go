package main

import (
	"bufio"
	"bytes"
	"fmt"
//	"strings"
	"github.com/nsf/termbox-go"
	"github.com/tatsushid/go-fastping"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

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
	addog(out, "test.txt")
	return out

}

func draw(i int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawLine(0, 0, "Press ESC to exit.")
	drawLine(2, 1, fmt.Sprintf("date: %v", time.Now()))
	if i > 0 {
		f, err := os.Open("test.txt")
		fatal(err)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		h := 2
		for scanner.Scan() {
			s := scanner.Text()
			drawLine(2, h, fmt.Sprintf("%v", s))
			if h == i {
				break
				}
			h++
			if err := scanner.Err(); err != nil {
			panic(err)
		}
			}
		drawLine(2, i+1, fmt.Sprintf("%v", Pinger("www.google.com")))
	}
	termbox.Flush()
}

func pollEvent() {
	var i int
	i = 0
	draw(i)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			default:
				i++
				draw(i)
			}
		default:
			draw(i)
		}
	}
}

func main() {
	ifExists("test.txt")
	defer os.Remove("test.txt")
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	pollEvent()
}
