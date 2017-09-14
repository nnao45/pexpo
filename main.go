package main

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/dariubs/percent"
	"github.com/nsf/termbox-go"
	"github.com/tatsushid/go-fastping"
	//	"io/ioutil"
	"log"
	"math"
	"net"
	"os"
	//	"regexp"
	"strconv"
	"strings"
	"time"
)

/*This global counter*/
var i int // "i" is all pings count.
var j int // "j" is all pings "per host" count.
var k int // "k" is scroll counter

/*This global buffer*/
var pbf bytes.Buffer // pbf is ping-list(textfile -> buffer).
var hbf bytes.Buffer // hbf is ping loss mapping to host.
var rbf bytes.Buffer // rbf is ping result list

const (
	DATE         = "2006 Jan 02 15:04:05.000Z07:00 JST"
	RED256       = 196
	BLUE256      = 21
	GREEN256     = 48
	WHITE256     = 255
	BENI256      = 13
	JUDGE_X      = 3
	HOST_X       = 7
	RTT_X        = 27
	DES_X        = 47
	LIST_H_X     = 70
	LIST_P_X     = 90
	LIST_L_X     = 98
	ICMP_TIMEOUT = 3
	DRAW_UP_Y    = 3
	DRAW_DW_Y    = 2
)

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

func drawLineColor(x, y int, str string, code int) {
	termbox.SetOutputMode(termbox.Output256)
	color := termbox.Attribute(code + 1)
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func drawLineColorful(x, y int, str string, strcode int, backcode int) {
	termbox.SetOutputMode(termbox.Output256)
	color := termbox.Attribute(strcode + 1)
	backgroundColor := termbox.Attribute(backcode + 1)
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

//func Pinger(host string, index int) (s string, flag string) {
func Pinger(host string, index int) (s string) {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", host)
	if err != nil {
		termbox.Close()
		panic(err)
	}
	p.AddIPAddr(ra)

	var out string
	var res string
	receiver := make(chan string, 100000)
	go func() {
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			//out = "Host: " + host + " IP Addr: " + addr.String() + " receive, RTT: " + rtt.String() + "\n"
			//out = "Host: " + host + " receive, RTT: " + rtt.String() + "\n"
			//out = host + " " + rtt.String() + "\n"
			out = host + " " + rtt.String()
			receiver <- out
		}
	}()
	p.OnIdle = func() {
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}

	timer := time.NewTimer(ICMP_TIMEOUT * time.Second)
	for {
		timer.Reset(ICMP_TIMEOUT * time.Second)
		select {
		case res = <-receiver:
			//return res, "o"
			res = "o " + res
			return res
		//case <-time.After(2 * time.Second):
		case <-timer.C:
			//res = "Host: " + host + " ping faild...\n"
			//res = "x " + host + " ping faild...\n"
			res = "x " + host + " ping...faild..."
			fres := strconv.Itoa(index) + "\n"
			hbf.WriteString(fres)
			//return res, "x"
			return res
		}
	}
}

func drawLoop() {
	for {
		j++
		//termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		drawHostList()
		var maxX int
		var maxY int
		index := 3
		maxX, maxY = termbox.Size()
		drawLine(maxX-27, 0, "Press Esc ,Ctrl+C to exit.")
		//_, maxY = getTermSize()
		//drawRed(20, 0, fmt.Sprintf("%v:%v", maxX, maxY))
		//drawRed(50, 0, fmt.Sprintf("%v", maxY))

		killKey := make(chan termbox.Key)
		//		resizeTerm := make(chan bool)
		go keyEventLoop(killKey)
		//		go getTermSize(resizeTerm)
		go func() {
			for {
				select {
				case wait := <-killKey:
					switch wait {
					case termbox.KeyEsc, termbox.KeyCtrlC:
						termbox.Close()
						os.Exit(0)
					}
					//				case <-resizeTerm:
					//					_, maxY = termbox.Size()
				}
			}
		}()

		pscanner := bufio.NewScanner(strings.NewReader(pbf.String()))
		for pscanner.Scan() {
			preps := pscanner.Text()
			preps_ary := strings.SplitN(preps, " ", 2)
			ps := preps_ary[0]
			des := preps_ary[1]
			res := Pinger(ps, index)
			res_ary := strings.SplitN(res, " ", 3)
			if maxY > i+DRAW_UP_Y+1 {
				drawFlag(JUDGE_X, i+DRAW_UP_Y, res_ary[0])
				drawFlag(JUDGE_X, 1, res_ary[0])
				if res_ary[0] == "o" {
					drawLine(HOST_X, i+DRAW_UP_Y, fmt.Sprintf("%v", res_ary[1]))
					drawLine(RTT_X, i+DRAW_UP_Y, fmt.Sprintf("%v", res_ary[2]))
					drawLine(DES_X, i+DRAW_UP_Y, fmt.Sprintf("%v", des))
				} else if res_ary[0] == "x" {
					drawLineColor(HOST_X, i+DRAW_UP_Y, fmt.Sprintf("%v", res_ary[1]), RED256)
					drawLineColor(RTT_X, i+DRAW_UP_Y, fmt.Sprintf("%v", res_ary[2]), RED256)
					drawLineColor(DES_X, i+DRAW_UP_Y, fmt.Sprintf("%v", des), RED256)
				}
			} else {
				/*ping-list clear*/
				fill(HOST_X, 3, 63, maxY-4, termbox.Cell{Ch: ' '})

				drawFlag(JUDGE_X, maxY-DRAW_DW_Y, res_ary[0])
				drawFlag(JUDGE_X, 1, res_ary[0])
				if res_ary[0] == "o" {
					drawLine(HOST_X, maxY-DRAW_DW_Y, fmt.Sprintf("%v", res_ary[1]))
					drawLine(RTT_X, maxY-DRAW_DW_Y, fmt.Sprintf("%v", res_ary[2]))
					drawLine(DES_X, maxY-DRAW_DW_Y, fmt.Sprintf("%v", des))
				} else if res_ary[0] == "x" {
					drawLineColor(HOST_X, maxY-DRAW_DW_Y, fmt.Sprintf("%v", res_ary[1]), RED256)
					drawLineColor(RTT_X, maxY-DRAW_DW_Y, fmt.Sprintf("%v", res_ary[2]), RED256)
					drawLineColor(DES_X, maxY-DRAW_DW_Y, fmt.Sprintf("%v", des), RED256)
				}
				var rc int
				rc = rc - k
				rscanner := bufio.NewScanner(strings.NewReader(rbf.String()))
				for rscanner.Scan() {
					rs := rscanner.Text()
					if rc > 0 {
						rs_ary := strings.SplitN(rs, " ", 4)
						drawFlag(JUDGE_X, rc+2, rs_ary[0])
						if rs_ary[0] == "o" {
							drawLine(HOST_X, rc+2, fmt.Sprintf("%v", rs_ary[1]))
							drawLine(RTT_X, rc+2, fmt.Sprintf("%v", rs_ary[2]))
							drawLine(DES_X, rc+2, fmt.Sprintf("%v", rs_ary[3]))
						} else if rs_ary[0] == "x" {
							drawLineColor(HOST_X, rc+2, fmt.Sprintf("%v", rs_ary[1]), RED256)
							drawLineColor(RTT_X, rc+2, fmt.Sprintf("%v", rs_ary[2]), RED256)
							drawLineColor(DES_X, rc+2, fmt.Sprintf("%v", rs_ary[3]), RED256)
						}
					} else {
						rs = ""
					}
					rc++
				}
				k++
			}
			pres := res_ary[0] + " " + res_ary[1] + " " + res_ary[2] + " " + des + "\n"
			rbf.WriteString(pres)
			drawLineColor(LIST_P_X, index, fmt.Sprintf("%.2f", Round(percent.PercentOf(drawLoss(index), j), 2)), GREEN256)
			drawLineColor(LIST_L_X, index, fmt.Sprintf("%v loss", drawLoss(index)), GREEN256)
			drawLine(HOST_X, 1, fmt.Sprintf("%v", "Host"))
			drawLine(RTT_X, 1, fmt.Sprintf("%v", "Response"))
			drawLine(DES_X, 1, fmt.Sprintf("%v", "Description"))
			fill(1, 0, 64, 1, termbox.Cell{Ch: '='})
			fill(1, 2, 64, 1, termbox.Cell{Ch: '='})
			fill(1, maxY-1, 64, 1, termbox.Cell{Ch: '='})
			fill(JUDGE_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(JUDGE_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
			fill(HOST_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(HOST_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
			fill(RTT_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(RTT_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
			fill(DES_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(DES_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
			fill(64, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(64, 1, 1, 1, termbox.Cell{Ch: '|'})
			//t := time.Now()
			//drawLine(2, 1, fmt.Sprintf("date: %v", t.Format(DATE)))
			//drawLine(2, 1, fmt.Sprintf("date: %v", t.Format(DATE)))
			termbox.Flush()
			i++
			index++
			if err := pscanner.Err(); err != nil {
				panic(err)
			}
		}
	}
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func drawFlag(x int, y int, flag string) {
	if flag == "o" {
		drawLineColor(x, y, fmt.Sprintf("%v", flag), BLUE256)
	} else if flag == "x" {
		drawLineColor(x, y, fmt.Sprintf("%v", flag), RED256)
	}
}

func drawHostList() {
	hi := 3
	//drawLineColor(LIST_H_X, 1, fmt.Sprintf("%v", "Loss counter Per host."), GREEN256)
	drawLineColorful(LIST_H_X-1, 1, fmt.Sprintf("%v", "      Now, Loss counting Per host.     "), WHITE256, BENI256)
	drawLine(LIST_H_X, 2, fmt.Sprintf("%v", "Hostname"))
	drawLine(LIST_P_X, 2, fmt.Sprintf("%v", "Loss(%)"))
	drawLine(LIST_L_X, 2, fmt.Sprintf("%v", "Loss(sum)"))
	scanner := bufio.NewScanner(strings.NewReader(pbf.String()))
	for scanner.Scan() {
		pres := scanner.Text()
		pres_ary := strings.SplitN(pres, " ", 2)
		s := pres_ary[0]
		drawLineColor(LIST_H_X, hi, fmt.Sprintf("%v", s), GREEN256)
		if j <= 1 {
			drawLineColor(LIST_P_X, hi, fmt.Sprintf("%v", "0.000"), GREEN256)
			drawLineColor(LIST_L_X, hi, fmt.Sprintf("%v", "0 loss"), GREEN256)
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

func getTermSize(resizeTerm chan bool) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			resizeTerm <- true
		default:
		}
	}
}

func keyEventLoop(killKey chan termbox.Key) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			killKey <- ev.Key
		//case termbox.EventResize:
		//			layout.termW, layout.termH = termbox.Size()
		//			drawHeader()
		default:
		}
	}
}

func init() {
	pl, err := os.Open("ping-list")
	fatal(err)
	defer pl.Close()
	scanner := bufio.NewScanner(pl)
	for scanner.Scan() {
		s := scanner.Text()
		if !strings.HasPrefix(s, "#") {
			for {
				if strings.HasPrefix(s, " ") {
					s_ary := strings.SplitN(s, " ", 2)
					s = s_ary[1]
				} else {
					break
				}
			}
			if !strings.Contains(s, " ") {
				s = s + " noname_host"
			} else {
				s_ary := strings.SplitN(s, " ", 2)
				s_ary[1] = strings.TrimSpace(s_ary[1])
				s = s_ary[0] + " " + s_ary[1]
			}
			s = s + "\n"
		} else {
			s = ""
		}
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
	drawLoop()
}
