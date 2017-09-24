/*

###################
pexpo Demo Console.
console 130x45
###################

=================================================================                    Ctrl+S: Stop & Restart, Esc or Ctrl+C: Exit.
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
//	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"os"
//	"os/exec"
	"os/user"
	"path/filepath"

	"strconv"
	"strings"
	"time"

	"github.com/dariubs/percent"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"github.com/tatsushid/go-fastping"
)

/*This global flag*/
var timeout = flag.Duration("t", time.Second*ICMP_TIMEOUT, "")
var interval = flag.Duration("i", time.Millisecond*ICMP_INTERVAL, "")
var pinglist = flag.String("f", PING_LIST, "")
var arp_entries = flag.Bool("A", false, "")
var httping = flag.Bool("H", false, "")
var sslping = flag.Bool("S", false, "")
var editor = flag.Bool("E", true, "")

/*This Used by func flag.Usage()*/
var usage = `
Usage:
    pexpo | pexpo.exe [-i interval] [-t timeout] [-f ping-list] [-A] [-H] [-S]

Examples:
    ./pexpo -i 500ms -t 1s -f /usr/local/ping-list.txt
    pexpo.exe -i 500ms -t 1s -f C:\Users\arale\Desktop\ping-list.txt

Option:
    -i Sending ICMP interval time(Default:500ms, should not be lower this).
       You must not use "200" or "1" or..., must use "200ms" or "1s" or ... , so use with time's unit.

    -t Sending ICMP timeout time(Default:3s)
       You must not use "200" or "1" or..., must use "200ms" or "1s" or ... , so use with time's unit.
       this "timeout" is Exact meaning, fastping.NewPinger() receives OnRecv struct value interval.

    -f Using Filepath of ping-list(Default:current_dir/ping-list.txt).

    -A If you want to write on ping-list -- such as Cisco's show ip arp -- , 
       "Internet  10.0.0.1                0   ca01.18cc.0038  ARPA   Ethernet2/0",
       Ignoring string "Internet", So It is good as you copy&paste show ip arp line.

<HTTP mode options!>

Examples:
    ./pexpo -H -i 500ms -t 1s -f /usr/local/curl-list.txt
    pexpo.exe -S -i 500ms -t 1s -f C:\Users\arale\Desktop\curl-list.txt
	(If you want to "Request, http and https", Using Both -H & -S.)
	   
	-H This optison is like "curl". So you Sending HTTP(:80) GET Request instead of the PING...!
	   
	-S This optison is like "curl". So you Sending HTTP"S"(:443) GET Request instead of the PING...!
	
	-H or -S options HTTP/HTTPS GET Request instead of the PING.
	(Just like, curl -LIs www.google.com -o /dev/null -w '%{http_code}\n')
	This Request is ververy simple GET Request, Only Getting status code(No header, No form, No getting data.)

	And, if http status code is "200", string color is Blue, else Red.
`

const (
	/*Used by logging*/
	DAY        = "20060102"
	DATE       = "2006-01-02 15:04:05.000"
	RESULT_DIR = ".pexpo"

	/*Default ping-list*/
	PING_LIST = "ping-list.txt"

	/*This values disigning terminal*/
	COLUMN = 17
	//JUDGE_X   = 3
	HOST_X    = 7
	RTT_X     = 27
	DES_X     = 47
	START_X   = 1
	EDGE_X    = 65
	LIST_H_X  = 70
	LIST_P_X  = 90
	LIST_L_X  = 100
	LIST_D_X  = 110
	DRAW_UP_Y = 3
	DRAW_DW_Y = 2

	/*Sending ICMP Param*/
	ICMP_INTERVAL = 500
	ICMP_TIMEOUT  = 3

	/*pexpo's version*/
	VERSION = "1.22"
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

func keyEventLoop(killKey chan termbox.Key) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			killKey <- ev.Key
		default:
		}
	}
}

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for n := 0; n < len(runes); n += 1 {
		termbox.SetCell(x+n, y, runes[n], color, backgroundColor)
	}
}

func drawLineColor(x, y int, str string, code termbox.Attribute) {
	termbox.SetOutputMode(termbox.Output256)
	color := code
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for n := 0; n < len(runes); n += 1 {
		termbox.SetCell(x+n, y, runes[n], color, backgroundColor)
	}
}

func drawLineColorful(x, y int, str string, strcode, backcode termbox.Attribute) {
	termbox.SetOutputMode(termbox.Output256)
	color := strcode
	backgroundColor := backcode
	runes := []rune(str)

	for n := 0; n < len(runes); n += 1 {
		termbox.SetCell(x+n, y, runes[n], color, backgroundColor)
	}
}

func drawFlag(x, y int, flag string) {
	if flag == "o" || flag == "200" {
		drawLineColor(x, y, fmt.Sprintf("%v", flag), termbox.ColorBlue)
	} else {
		drawLineColor(x, y, fmt.Sprintf("%v", flag), termbox.ColorRed)
	}
}

func drawSeq(hx, rx, dx, y int, flag, r1, r2, des string) {
	if flag == "o" || flag == "200" {
		drawLine(hx, y, fmt.Sprintf("%v", runewidth.Truncate(r1, COLUMN, "")))
		drawLine(rx, y, fmt.Sprintf("%v", runewidth.Truncate(r2, COLUMN, "")))
		drawLine(dx, y, fmt.Sprintf("%v", runewidth.Truncate(des, COLUMN, "")))
	} else {
		drawLineColor(hx, y, fmt.Sprintf("%v", runewidth.Truncate(r1, COLUMN, "")), termbox.ColorRed)
		drawLineColor(rx, y, fmt.Sprintf("%v", runewidth.Truncate(r2, COLUMN, "")), termbox.ColorRed)
		drawLineColor(dx, y, fmt.Sprintf("%v", runewidth.Truncate(des, COLUMN, "")), termbox.ColorRed)
	}

}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

/*This Core of the sendig ICMP engine*/
func Pinger(host string) []string {
	p := fastping.NewPinger()

	/*Selecting IPv4 or IPv6*/
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

	p.MaxRTT = *interval
	var out []string
	var res []string
	receiver := make(chan []string, 3)

	/*Received value from fastping.NewPinger()*/
	go func() {
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			out = append(out, host, rtt.String())
			receiver <- out
			defer close(receiver)
		}
	}()
	p.OnIdle = func() {
	}
	err = p.Run()
	/*if err != nil {
	                fmt.Println(err)
			        }*/
	fatal(err)

	/*Set the timeout timer*/
	timer := time.NewTimer(*timeout)
	for {
		timer.Reset(*timeout)
		select {
		case res = <-receiver:
			res = append([]string{"o"}, res...)
			return res
		case <-timer.C:
			res = append(res, "x", host, "ping...faild...")
			return res
		}
	}
}

func curlCheck(url string) []string {
	var out []string
	var res []string
	receiver := make(chan []string, 3)
	done := make(chan struct{}, 0)
	if *httping && *sslping {
		if !strings.Contains(url, "https://") && !strings.Contains(url, "http://") {
			url = "https://" + url
		}
	} else {
		if *sslping {
			if !strings.Contains(url, "https://") {
				url = "https://" + url
			}
		} else if *httping {
			if !strings.Contains(url, "http://") {
				url = "http://" + url
			}
		}
	}
	time_start := time.Now()
	c_timeout := time.Duration(*timeout * time.Second)
	go func() {
		client := http.Client{
			Timeout: c_timeout,
		}
		resp, err := client.Get(url)
		if err != nil {
			<-done
			defer close(done)
		}
		out = append(out, strconv.Itoa(resp.StatusCode), url, time.Since(time_start).String())
		receiver <- out
		defer close(receiver)

		defer resp.Body.Close()
	}()

	timer := time.NewTimer(*timeout)
	for {
		timer.Reset(*timeout)
		select {
		case res = <-receiver:
			return res
		case <-timer.C:
			if *sslping {
				res = append(res, "000", url, "ssl...no_response")
			} else {
				res = append(res, "000", url, " http...no_response")
			}
			return res
		case <-done:
			if *sslping {
				res = append(res, "000", url, "ssl...no_response")
			} else {
				res = append(res, "000", url, "http...no_response")
			}
			return res
		}
	}
}

/*This is Main loop*/
func drawLoop(maxX, maxY int, stop, restart, received chan struct{}) {

	var i int // "i" is all pings count.
	var j int // "j" is all pings "per host" count.
	var k int // "k" is scroll counter

	var pbf_ary []string // pbf_ary is ping-list(textfile -> buffer).
	var rbf_ary []string // rbf_ary is ping result list.
	var hbf_ary []string // hbf_ary is ping loss counter map to per host.

	/*1st key loop lock open*/
	received <- struct{}{}

	var JUDGE_X int
	if *httping || *sslping {
		JUDGE_X = 2
	} else {
		JUDGE_X = 3
	}

	for {
		/*Counting per running This function*/
		j++

		/*Userd by hbf, index must be Reinitialize in per loop*/
		index := DRAW_UP_Y

		/*This Aciton, Only 1st loop!!*/
		if j <= 1 {
			drawLine(maxX-44, 0, "Ctrl+S: Stop & Restart, Esc or Ctrl+C: Exit.")
			drawLine(maxX-9, maxY-1, fmt.Sprintf("ver. %v", VERSION))
			drawLineColorful(LIST_H_X-1, 1, fmt.Sprintf("%v", "           Now, Loss counting Per host.            "), termbox.ColorDefault, termbox.ColorMagenta)
			drawLineColor(LIST_H_X, 2, fmt.Sprintf("%v", "Hostname"), termbox.ColorWhite)
			drawLineColor(LIST_P_X, 2, fmt.Sprintf("%v", "Loss(%)"), termbox.ColorWhite)
			drawLineColor(LIST_L_X, 2, fmt.Sprintf("%v", "Loss(sum)"), termbox.ColorWhite)
			drawLineColor(LIST_D_X, 2, fmt.Sprintf("%v", "Dead Now?"), termbox.ColorWhite)
			drawLine(HOST_X, 1, fmt.Sprintf("%v", "Host"))
			drawLine(RTT_X, 1, fmt.Sprintf("%v", "Response"))
			drawLine(DES_X, 1, fmt.Sprintf("%v", "Description"))
			fill(START_X, 0, EDGE_X, 1, termbox.Cell{Ch: '='})
			fill(START_X, 2, EDGE_X, 1, termbox.Cell{Ch: '='})
			fill(START_X, maxY-1, EDGE_X, 1, termbox.Cell{Ch: '='})
			fill(START_X, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(START_X, 1, 1, 1, termbox.Cell{Ch: '|'})
			fill(HOST_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(HOST_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
			fill(RTT_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(RTT_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
			fill(DES_X-2, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(DES_X-2, 1, 1, 1, termbox.Cell{Ch: '|'})
			fill(EDGE_X, 3, 1, maxY-4, termbox.Cell{Ch: '|'})
			fill(EDGE_X, 1, 1, 1, termbox.Cell{Ch: '|'})

			pl, err := os.Open(*pinglist)
			if err != nil {
				termbox.Close()
				panic(err)
			}
			defer pl.Close()
			plscanner := bufio.NewScanner(pl)
			for plscanner.Scan() {
				s := plscanner.Text()

				/*If blank line, transform to "#"*/
				if s == "" {
					s = "#" + s
				}

				/*# is comment out line*/
				if !strings.HasPrefix(s, "#") {

					/*Deleting "\t"...tab key*/
					for {
						if strings.Contains(s, "\t") {
							s_ary := strings.SplitN(s, "\t", 2)
							s = s_ary[0] + " " + s_ary[1]
						} else {
							break
						}
					}
					/*Deleting consecutive white space "head"
					  "                 8.8.8.8 google.com" */
					for {
						if strings.HasPrefix(s, " ") {
							s_ary := strings.SplitN(s, " ", 2)
							s = s_ary[1]
						} else {
							break
						}
					}

					/*No description, Put in "noname_host"
					  If not this statement, After, Will read blank array,
					  So, occuring panic error*/
					if !strings.Contains(s, " ") {
						s = s + " noname_host"
					} else {

						/*For -a option
						Ignoring string, "Internet"*/
						if *arp_entries && strings.HasPrefix(s, "Internet") {
							s_ary := strings.SplitN(s, "  ", 2)
							s = s_ary[1]
						}
						/*Deleting consecutive white space "between"
						  "8.8.8.8                          google.com" */
						s_ary := strings.SplitN(s, " ", 2)
						s_ary[1] = strings.TrimSpace(s_ary[1])
						s = s_ary[0] + " " + s_ary[1]
					}

					if !*httping || !*sslping {
						s_ary := strings.SplitN(s, " ", 2)
						if *httping && strings.Contains(s_ary[0], "https://") {
							termbox.Close()
							fmt.Printf("Sorry, %v is not http protocol.\n", s_ary[0])
							fmt.Printf("Please, Check your %v.\n", *pinglist)
							os.Exit(1)
						} else if *sslping && strings.Contains(s_ary[0], "http://") {
							termbox.Close()
							fmt.Printf("Sorry, %v is not https protocol.\n", s_ary[0])
							fmt.Printf("Please, Check your %v.\n", *pinglist)
							os.Exit(1)
						}
					}

				/*ping-list -> pbf*/
				pbf_ary = append(pbf_ary, s)

				} 
				
				if err := plscanner.Err(); err != nil {
					panic(err)
				}
			}
			for n, pres := range pbf_ary {
				s_ary := strings.SplitN(pres, " ", 2)
				s := s_ary[0]
				drawLineColor(LIST_H_X, n+DRAW_UP_Y, fmt.Sprintf("%v", runewidth.Truncate(s, COLUMN, "")), termbox.ColorGreen)
				drawLineColor(LIST_P_X, n+DRAW_UP_Y, fmt.Sprintf("%v", "0.00"), termbox.ColorGreen)
				drawLineColor(LIST_L_X, n+DRAW_UP_Y, fmt.Sprintf("%v", "0   loss"), termbox.ColorGreen)
			}
		}

		/*Reading Ping-list per line*/
		for _, preps := range pbf_ary {

			/*For Stop & Restart*/
			select {
			case <-stop:
				received <- struct{}{}
				<-restart
				received <- struct{}{}

			/*Default behavior*/
			default:
			}
			preps_ary := strings.SplitN(preps, " ", 2)
			ps := preps_ary[0]
			des := preps_ary[1]
			var res_ary []string
			if *httping || *sslping {
				time.Sleep(*interval)
				res_ary = curlCheck(ps)
			} else {
				res_ary = Pinger(ps)
			}
			if res_ary[0] != "o" && res_ary[0] != "200" {
				hbf_ary = append(hbf_ary, res_ary[1])
			}
			/*Before Scrolling To the bottom*/
			if maxY > i+DRAW_UP_Y+1 {
				drawFlag(JUDGE_X, i+DRAW_UP_Y, res_ary[0])
				drawFlag(JUDGE_X, 1, res_ary[0])
				drawSeq(HOST_X, RTT_X, DES_X, i+DRAW_UP_Y, res_ary[0], res_ary[1], res_ary[2], des)

				/*After Scrolling To the bottom*/
			} else {
				/*ping-list clear*/
				fill(JUDGE_X+1, DRAW_UP_Y, 1, maxY-4, termbox.Cell{Ch: ' '})
				fill(HOST_X+1, DRAW_UP_Y, COLUMN-1, maxY-4, termbox.Cell{Ch: ' '})
				fill(RTT_X+1, DRAW_UP_Y, COLUMN-1, maxY-4, termbox.Cell{Ch: ' '})
				fill(DES_X+1, DRAW_UP_Y, COLUMN-1, maxY-4, termbox.Cell{Ch: ' '})

				drawFlag(JUDGE_X, maxY-DRAW_DW_Y, res_ary[0])
				drawFlag(JUDGE_X, 1, res_ary[0])
				drawSeq(HOST_X, RTT_X, DES_X, maxY-DRAW_DW_Y, res_ary[0], res_ary[1], res_ary[2], des)
				/*rc is count Reading rbf After Scrolling To the bottom*/
				var rc int
				var tmp_ary []string
				/*"rc" -"k" -> "All Result" - "Line of Don't want to see" */
				rc = rc - k
				for _, rs := range rbf_ary {
					if rc > 0 {
						rs_ary := strings.SplitN(rs, " ", 4)
						drawFlag(JUDGE_X, rc+2, rs_ary[0])
						drawSeq(HOST_X, RTT_X, DES_X, rc+2, rs_ary[0], rs_ary[1], rs_ary[2], rs_ary[3])
						tmp_ary = append(tmp_ary, rs)
					} 
					rc++
				}
				k++
				copy(tmp_ary, rbf_ary)
			}
			/*finish Reading Ping-list & Drawing Result.
			  After, Logging, & Drawing Loss Counter*/

			var pres []string
			pres = append(pres, res_ary[0], res_ary[1], res_ary[2], des)
			
			/*Logging rbf -> This buffer Called by Next Drawing*/
			rbf_ary = append(rbf_ary, strings.Join(pres, " "))

			/*Logging All Result with time stamp*/
			day := time.Now()
			date := time.Now()
			formating_day := day.Format(DAY)
			formating_date := date.Format(DATE)
			log := "[" + formating_date + "]" + " " + strings.Join(pres, " ")
			result := "result_" + formating_day + ".txt"
			u, err := user.Current()
			fatal(err)
			rfile := filepath.Join(u.HomeDir, RESULT_DIR, result)
			addog(log, rfile)

			/*Refreshing Loss Counter*/
			fill(LIST_P_X, index, 10, 1, termbox.Cell{Ch: ' '})

			/*Loss counting from hbf*/
			var c int
			for _, loss := range hbf_ary {
				/*So, If pexpo had been sending ICMP loss, pexpo logging per host to the hbf
				This func loss counting*/
				if loss == res_ary[1] {
					c++
				}
			}

			/*Drawing Loss Counter*/
			drawLineColor(LIST_P_X, index, fmt.Sprintf("%.2f", Round(percent.PercentOf(c, j), 2)), termbox.ColorGreen)
			drawLineColor(LIST_L_X, index, fmt.Sprintf("%v", c), termbox.ColorGreen)
			drawLineColor(LIST_L_X+4, index, fmt.Sprintf("%v", "loss"), termbox.ColorGreen)

			/*Drawing the Dead stamp*/
			if res_ary[0] == "o" || res_ary[0] == "200" {
				fill(LIST_D_X, index, 9, 1, termbox.Cell{Ch: ' '})
			} else {
				drawLineColor(LIST_D_X, index, fmt.Sprintf("%v", "Dead Now!"), termbox.ColorRed)
			}

			/*Drawing Done*/
			termbox.Flush()

			/*All couting per sending ICMP*/
			i++

			/*"index" for the mapping host to the Loss counter*/
			index++
		}
	}
}

func init() {
	flag.Usage = func() {
		fmt.Printf(usage)
	}

	flag.Parse()

	u, err := user.Current()
	fatal(err)
	rdir := filepath.Join(u.HomeDir, RESULT_DIR)
	err = os.MkdirAll(rdir, 0755)
	fatal(err)
	
	/*
	if *editor {
	
	edited := make(chan struct{}, 0)
		go func () {
		cmd := exec.Command(`C:\Program Files\Notepad++\notepad++.exe`, *pinglist)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil{
			panic(err)
			}
		edited <- struct{}{}
		}()
		<- edited
	}
	*/
}

func main() {

	/*termbox start*/
	err := termbox.Init()
	fatal(err)

	defer termbox.Close()

	maxX, maxY := termbox.Size()

	/*stop channel is for stopping drawLoop()*/
	stop := make(chan struct{}, 0)

	/*stop channel is for restarting drawLoop()*/
	restart := make(chan struct{}, 0)

	/*received channel is received message from drawLoop()*/
	received := make(chan struct{}, 0)

	/*killKey channel is received HW key interrupt*/
	killKey := make(chan termbox.Key)

	/*sleep flag*/
	sleep := false

	go keyEventLoop(killKey)
	go drawLoop(maxX, maxY, stop, restart, received)

loop:
	<-received
	for {
		select {
		case wait := <-killKey:
			switch wait {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				return
			case termbox.KeyCtrlS:
				if sleep == false {
					fill(maxX-44, 0, 45, 1, termbox.Cell{Ch: ' '})
					drawLineColor(maxX-48, 0, "Stop Now!! Crtl+S: Restart, Esc or Ctrl+C: Exit.", termbox.ColorYellow)
					stop <- struct{}{}
					sleep = true
					goto loop
				} else if sleep == true {
					fill(maxX-48, 0, 49, 1, termbox.Cell{Ch: ' '})
					drawLine(maxX-44, 0, "Ctrl+S: Stop & Restart, Esc or Ctrl+C: Exit.")
					restart <- struct{}{}
					sleep = false
					goto loop
				}
			}
		}
	}
}
