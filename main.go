/*

###################
pexpo Demo Console.
console 130x45
###################

 =================================================================                                           Esc ,Ctrl+C to exit.
 | o | Host              | Response          | Description       |              Now, Loss counting Per host.            
 =================================================================    Hostname            Loss(%)   Loss(sum) Dead Now?
 | o | 77.88.8.3         | 286.120512ms      | Yandex.DNS        |    www.yahoo.com       0.000     0   loss
 | o | 180.76.76.76      | 37.444591ms       | Baidu DNS         |    192.168.1.201       100.00    10  loss  Dead Now!
 | o | 114.114.114.114   | 288.835931ms      | Baidu DNS         |    2001:4860:4860:...  0.000     0   loss
 | o | 80.80.80.80       | 74.98565ms        | Freenom World     |    8.8.8.8             0.000     0   loss
 | o | 80.80.81.81       | 73.907249ms       | Freenom World     |    8.8.4.4             0.000     0   loss
 | o | 8.26.56.26        | 213.717929ms      | Comodo Secure DNS |    208.67.222.123      0.000     0   loss
 | o | 8.20.247.20       | 209.27997ms       | Comodo Secure DNS |    208.67.220.123      0.000     0   loss
 | o | 106.186.17.181    | 74.254678ms       | OpenNIC           |    216.146.35.35       0.000     0   loss
 | o | 106.185.41.36     | 66.059639ms       | OpenNIC           |    216.146.36.36       0.000     0   loss
 | o | www.yahoo.com     | 75.149983ms       | Yahoo!!!          |    77.88.8.8           0.000     0   loss
 | x | 192.168.1.201     | ping...faild...   | Host is Dead!     |    77.88.8.1           0.000     0   loss
 | o | 2001:4860:4860:...| 1.197034ms        | Google_IPv6       |    77.88.8.88          0.000     0   loss
 | o | 8.8.8.8           | 941.56µs          | nandedaaaaaaaaa...|    77.88.8.2           0.000     0   loss
 | o | 8.8.4.4           | 991.28µs          | Google_IPv4       |    77.88.8.7           0.000     0   loss
 | o | 208.67.222.123    | 944.168µs         | OpenDNS           |    77.88.8.3           0.000     0   loss
 | o | 208.67.220.123    | 948.797µs         | OpenDNS           |    180.76.76.76        0.000     0   loss
 | o | 216.146.35.35     | 954.225µs         | Dyn Internet Guide|    114.114.114.114     0.000     0   loss
 | o | 216.146.36.36     | 240.103655ms      | Dyn Internet Guide|    80.80.80.80         0.000     0   loss
 | o | 77.88.8.8         | 286.464099ms      | Yandex.DNS        |    80.80.81.81         0.000     0   loss
 | o | 77.88.8.1         | 358.621804ms      | Yandex.DNS        |    8.26.56.26          0.000     0   loss
 | o | 77.88.8.88        | 358.995054ms      | Yandex.DNS        |    8.20.247.20         0.000     0   loss
 | o | 77.88.8.2         | 287.493195ms      | Yandex.DNS        |    106.186.17.181      0.000     0   loss
 | o | 77.88.8.7         | 358.558967ms      | Yandex.DNS        |    106.185.41.36       0.000     0   loss
 | o | 77.88.8.3         | 285.775403ms      | Yandex.DNS        |
 | o | 180.76.76.76      | 37.389211ms       | Baidu DNS         |
 | o | 114.114.114.114   | 288.735094ms      | Baidu DNS         |
 | o | 80.80.80.80       | 75.043223ms       | Freenom World     |
 | o | 80.80.81.81       | 73.908742ms       | Freenom World     |
 | o | 8.26.56.26        | 213.743171ms      | Comodo Secure DNS |
 | o | 8.20.247.20       | 209.260693ms      | Comodo Secure DNS |
 | o | 106.186.17.181    | 74.223363ms       | OpenNIC           |
 | o | 106.185.41.36     | 66.092089ms       | OpenNIC           |
 | o | www.yahoo.com     | 75.187195ms       | Yahoo!!!          |
 | x | 192.168.1.201     | ping...faild...   | Host is Dead!     |
 | o | 2001:4860:4860:...| 1.264779ms        | Google_IPv6       |
 | o | 8.8.8.8           | 977.278µs         | nandedaaaaaaaaa...|
 | o | 8.8.4.4           | 979.161µs         | Google_IPv4       |
 | o | 208.67.222.123    | 1.186372ms        | OpenDNS           |
 | o | 208.67.220.123    | 951.586µs         | OpenDNS           |
 | o | 216.146.35.35     | 930.594µs         | Dyn Internet Guide|
 | o | 216.146.36.36     | 240.093372ms      | Dyn Internet Guide|
 =================================================================
*/

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/dariubs/percent"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"github.com/tatsushid/go-fastping"
	"log"
	"math"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"path"
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

var timeout = flag.Duration("t", time.Second * ICMP_TIMEOUT, "")
var interval = flag.Duration("i", time.Millisecond * ICMP_INTERVAL, "")
var pinglist = flag.String("f", PING_LIST, "")
var arp_entries = flag.Bool("a", false, "")

var usage = `
Usage:
    pexpo | pexpo.exe [-i interval] [-t timeout] [-f ping-list] [-a arp_entries]

Examples:
    ./pexpo -i 500ms -t 1s -f /usr/local/ping-list
    pexpo.exe -i 500ms -t 1s -f C:\Users\arale\Desktop\ping-list

Option:
    -i Sending ICMP interval time(Default:500ms, should not be lower this).
       You must not use "200" or "1" or..., must use "200ms" or "1s" or ... , so use with time's unit.

    -t Sending ICMP timeout time(Default:3s)
       You must not use "200" or "1" or..., must use "200ms" or "1s" or ... , so use with time's unit.
       this "timeout" is Exact meaning, fastping.NewPinger() receives OnRecv struct value interval.

    -f Using Filepath of ping-list(Default:current_dir/ping-list.txt).

    -a If you want to write on ping-list -- such as Cisco's show ip arp -- , 
       "Internet  10.0.0.1                0   ca01.18cc.0038  ARPA   Ethernet2/0",
       Ignoring string "Internet", So It is good as you copy&paste show ip arp line.
`


const (
	DAY           = "20060102"
	DATE          = "2006-01-02 15:04:05.000"
	PING_LIST     = "ping-list.txt"
	RESULT_DIR    = ".pexpo"
	RED256        = 196
	BLUE256       = 21
	GREEN256      = 48
	WHITE256      = 255
	BENI256       = 13
	COLUMN        = 18
	JUDGE_X       = 3
	HOST_X        = 7
	RTT_X         = 27
	DES_X         = 47
	START_X       = 1
	EDGE_X        = 65
	LIST_H_X      = 70
	LIST_P_X      = 90
	LIST_L_X      = 100
	LIST_D_X      = 110
	ICMP_INTERVAL = 500
	ICMP_TIMEOUT  = 3
	DRAW_UP_Y     = 3
	DRAW_DW_Y     = 2
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

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for n := 0; n < len(runes); n += 1 {
		termbox.SetCell(x+n, y, runes[n], color, backgroundColor)
	}
}

//func drawLineColor(x, y int, str string, code int) {
func drawLineColor(x, y int, str string, code termbox.Attribute) {
	termbox.SetOutputMode(termbox.Output256)
	//color := termbox.Attribute(code + 1)
	color := code
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for n := 0; n < len(runes); n += 1 {
		termbox.SetCell(x+n, y, runes[n], color, backgroundColor)
	}
}

//func drawLineColorful(x, y int, str string, strcode int, backcode int) {
func drawLineColorful(x, y int, str string, strcode, backcode termbox.Attribute) {
	termbox.SetOutputMode(termbox.Output256)
	//color := termbox.Attribute(strcode + 1)
	color := strcode
	//backgroundColor := termbox.Attribute(backcode + 1)
	backgroundColor := backcode
	runes := []rune(str)

	for n := 0; n < len(runes); n += 1 {
		termbox.SetCell(x+n, y, runes[n], color, backgroundColor)
	}
}

func Pinger(host string, index int) (s string) {
	p := fastping.NewPinger()
	netProto := "ip4:icmp"
	if strings.Index(host, ":") != -1 {
		netProto = "ip6:ipv6-icmp"
	}
	ra, err := net.ResolveIPAddr(netProto, host)
	if err != nil {
		termbox.Close()
		panic(err)
	}
	p.AddIPAddr(ra)


	//p.MaxRTT = time.Millisecond * ICMP_INTERVAL
	p.MaxRTT = *interval
	var out string
	var res string
	receiver := make(chan string, EDGE_X)
	go func() {
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			//out = "Host: " + host + " IP Addr: " + addr.String() + " receive, RTT: " + rtt.String() + "\n"
			out = host + " " + rtt.String()
			receiver <- out
			defer close(receiver)
		}
	}()
	p.OnIdle = func() {
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}

	//timer := time.NewTimer(ICMP_TIMEOUT * time.Second)
	timer := time.NewTimer(*timeout)
	for {
		//timer.Reset(ICMP_TIMEOUT * time.Second)
		timer.Reset(*timeout)
		select {
		case res = <-receiver:
			res = "o " + res
			return res
		case <-timer.C:
			res = "x " + host + " ping...faild..."
			fres := strconv.Itoa(index) + "\n"
			hbf.WriteString(fres)
			return res
		}
	}
}

func drawLoop() {
	for {
		j++
		drawHostList()
		var maxX int
		var maxY int
		index := DRAW_UP_Y
		maxX, maxY = termbox.Size()
		drawLine(maxX-44, 0, "Ctrl+S: Stop & Restart, Esc or Ctrl+C: Exit.")
		//_, maxY = getTermSize()

		killKey := make(chan termbox.Key)
		sleep := false
		//stop := make(chan bool)
		restart := make(chan bool)
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
					case termbox.KeyCtrlS:
						if sleep == false {
						sleep = true
						} else {
						restart <- true
						sleep = false
						}
					}
					//				case <-resizeTerm:
					//					_, maxY = termbox.Size()
				}
			}
		}()

		pscanner := bufio.NewScanner(strings.NewReader(pbf.String()))
		for pscanner.Scan() {
		if sleep == true {
			<-restart
			continue
		} else {
			preps := pscanner.Text()
			preps_ary := strings.SplitN(preps, " ", 2)
			ps := preps_ary[0]
			des := preps_ary[1]
			res := Pinger(ps, index)
			res_ary := strings.SplitN(res, " ", 3)
			if maxY > i+DRAW_UP_Y+1 {
				drawFlag(JUDGE_X, i+DRAW_UP_Y, res_ary[0])
				drawFlag(JUDGE_X, 1, res_ary[0])
				drawSeq(HOST_X, RTT_X, DES_X, i+DRAW_UP_Y, res_ary[0], res_ary[1], res_ary[2], des)
			} else {
				/*ping-list clear*/
				fill(JUDGE_X+1, DRAW_UP_Y, 1, maxY-4, termbox.Cell{Ch: ' '})
				fill(HOST_X+1, DRAW_UP_Y, COLUMN-1, maxY-4, termbox.Cell{Ch: ' '})
				fill(RTT_X+1, DRAW_UP_Y, COLUMN-1, maxY-4, termbox.Cell{Ch: ' '})
				fill(DES_X+1, DRAW_UP_Y, COLUMN-1, maxY-4, termbox.Cell{Ch: ' '})

				drawFlag(JUDGE_X, maxY-DRAW_DW_Y, res_ary[0])
				drawFlag(JUDGE_X, 1, res_ary[0])
				drawSeq(HOST_X, RTT_X, DES_X, maxY-DRAW_DW_Y, res_ary[0], res_ary[1], res_ary[2], des)
				var rc int
				rc = rc - k
				rscanner := bufio.NewScanner(strings.NewReader(rbf.String()))
				for rscanner.Scan() {
					rs := rscanner.Text()
					if rc > 0 {
						rs_ary := strings.SplitN(rs, " ", 4)
						drawFlag(JUDGE_X, rc+2, rs_ary[0])
						drawSeq(HOST_X, RTT_X, DES_X, rc+2, rs_ary[0], rs_ary[1], rs_ary[2], rs_ary[3])
					} else {
						rs = ""
					}
					rc++
				}
				k++
			}
			pres := res_ary[0] + " " + res_ary[1] + " " + res_ary[2] + " " + des + "\n"
			logres := res_ary[0] + " " + res_ary[1] + " " + res_ary[2] + "\n"
			rbf.WriteString(pres)

			day  := time.Now()
			date := time.Now()
			formating_day := day.Format(DAY)
			formating_date := date.Format(DATE)
			log := "[" + formating_date + "]" + " " + logres
			result := "result_" + formating_day + ".txt"

			u, err := user.Current()
			if err != nil {
				fatal(err)
			}
			rfile := filepath.Join(u.HomeDir, RESULT_DIR, result)
			addog(log, rfile)

			//			drawLineColor(LIST_P_X, index, fmt.Sprintf("%.2f", Round(percent.PercentOf(drawLoss(index), j), 2)), GREEN256)
			//			drawLineColor(LIST_L_X, index, fmt.Sprintf("%v  loss", drawLoss(index)), GREEN256)
			drawLineColor(LIST_P_X, index, fmt.Sprintf("%.2f", Round(percent.PercentOf(drawLoss(index), j), 2)), termbox.ColorGreen)
			//drawLineColor(LIST_L_X, index, fmt.Sprintf("%v  loss", drawLoss(index)), termbox.ColorGreen)
			drawLineColor(LIST_L_X, index, fmt.Sprintf("%v", drawLoss(index)), termbox.ColorGreen)
			drawLineColor(LIST_L_X+4, index, fmt.Sprintf("%v", "loss"), termbox.ColorGreen)
			if res_ary[0] == "x" {
				//drawLineColor(LIST_D_X, index, fmt.Sprintf("%v", "Dead Now!"), RED256)
				drawLineColor(LIST_D_X, index, fmt.Sprintf("%v", "Dead Now!"), termbox.ColorRed)
			} else if res_ary[0] == "o" {
				fill(LIST_D_X, index, 9, 1, termbox.Cell{Ch: ' '})
			}
			if i <= 1 {
				drawLine(HOST_X, 1, fmt.Sprintf("%v", "Host"))
				drawLine(RTT_X, 1, fmt.Sprintf("%v", "Response"))
				drawLine(DES_X, 1, fmt.Sprintf("%v", "Description"))
				fill(START_X, 0, EDGE_X, 1, termbox.Cell{Ch: '='})
				fill(START_X, 2, EDGE_X, 1, termbox.Cell{Ch: '='})
				fill(START_X, maxY-1, EDGE_X, 1, termbox.Cell{Ch: '='})
				fill(JUDGE_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
				fill(JUDGE_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
				fill(HOST_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
				fill(HOST_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
				fill(RTT_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
				fill(RTT_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
				fill(DES_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
				fill(DES_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
				fill(EDGE_X, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
				fill(EDGE_X, 1, 1, 1, termbox.Cell{Ch: '|'})
			}
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
		drawLineColor(x, y, fmt.Sprintf("%v", flag), termbox.ColorBlue)
	} else if flag == "x" {
		drawLineColor(x, y, fmt.Sprintf("%v", flag), termbox.ColorRed)
	}
}

func drawSeq(hx, rx, dx, y int, flag, r1, r2, des string) {
	if flag == "o" {
		drawLine(hx, y, fmt.Sprintf("%v", runewidth.Truncate(r1, COLUMN, "...")))
		drawLine(rx, y, fmt.Sprintf("%v", runewidth.Truncate(r2, COLUMN, "...")))
		drawLine(dx, y, fmt.Sprintf("%v", runewidth.Truncate(des, COLUMN, "...")))
	} else if flag == "x" {
		drawLineColor(hx, y, fmt.Sprintf("%v", runewidth.Truncate(r1, COLUMN, "...")), termbox.ColorRed)
		drawLineColor(rx, y, fmt.Sprintf("%v", runewidth.Truncate(r2, COLUMN, "...")), termbox.ColorRed)
		drawLineColor(dx, y, fmt.Sprintf("%v", runewidth.Truncate(des, COLUMN, "...")), termbox.ColorRed)
	}

}

func drawHostList() {
	hi := 3
	if j <= 1 {
		drawLineColorful(LIST_H_X-1, 1, fmt.Sprintf("%v", "           Now, Loss counting Per host.            "), termbox.ColorDefault, termbox.ColorMagenta)
		drawLineColor(LIST_H_X, 2, fmt.Sprintf("%v", "Hostname"), termbox.ColorWhite)
		drawLineColor(LIST_P_X, 2, fmt.Sprintf("%v", "Loss(%)"), termbox.ColorWhite)
		drawLineColor(LIST_L_X, 2, fmt.Sprintf("%v", "Loss(sum)"), termbox.ColorWhite)
		drawLineColor(LIST_D_X, 2, fmt.Sprintf("%v", "Dead Now?"), termbox.ColorWhite)
	}
	scanner := bufio.NewScanner(strings.NewReader(pbf.String()))
	for scanner.Scan() {
		pres := scanner.Text()
		pres_ary := strings.SplitN(pres, " ", 2)
		s := pres_ary[0]
		drawLineColor(LIST_H_X, hi, fmt.Sprintf("%v", runewidth.Truncate(s, COLUMN, "...")), termbox.ColorGreen)
		if j <= 1 {
			drawLineColor(LIST_P_X, hi, fmt.Sprintf("%v", "0.000"), termbox.ColorGreen)
			drawLineColor(LIST_L_X, hi, fmt.Sprintf("%v", "0   loss"), termbox.ColorGreen)
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

/*
func getTermSize(resizeTerm chan bool) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			resizeTerm <- true
		default:
		}
	}
}
*/
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
	flag.Usage = func() {
				fmt.Printf(usage)
	}

	flag.Parse()

	u, err := user.Current()
	if err != nil {
		fatal(err)
	}
	rdir := filepath.Join(u.HomeDir, RESULT_DIR)
	err = os.MkdirAll(rdir, 0755)
	if err != nil {
		fatal(err)
	}
	//pl, err := os.Open(PING_LIST)
	pl, err := os.Open(path.Base(*pinglist))
	fatal(err)
	defer pl.Close()
	scanner := bufio.NewScanner(pl)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			s = "#" + s
		}
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
				if *arp_entries && strings.HasPrefix(s, "Internet") {
					s_ary := strings.SplitN(s, "  ", 2)
					s = s_ary[1]
				}
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
