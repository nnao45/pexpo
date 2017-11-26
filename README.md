[![v1.30](https://img.shields.io/badge/package-v1.30-ff69b4.svg)](https://github.com/nnao45/pexpo/releases/tag/1.30)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/nnao45/pexpo/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/nnao45/pexpo)](https://goreportcard.com/report/github.com/nnao45/pexpo)
[![Travis CI](https://travis-ci.org/nnao45/pexpo.svg?branch=master)](https://travis-ci.org/nnao45/pexpo)
[![platform](https://img.shields.io/conda/pn/conda-forge/python.svg?maxAge=2592000)]()
# pexpo
![result](https://user-images.githubusercontent.com/17565502/30773031-041851a6-a0a3-11e7-90be-81199aa12676.png)  
pexpo is ping sending tui tool with cool clomun & logging loss-count in the multi platforms(Windows, Mac, Linux...).  
pexpo has tui engine is [termbox-go](https://github.com/nsf/termbox-go), sending ICMP engine is [go-fastping](https://github.com/tatsushid/go-fastping).  
So, pexpo's code is NATIVE [The Go Programming Language](http://golang.org) application. 
This is inspired the [pinger](https://github.com/hirose31/pinger), [Exping](http://www.woodybells.com/exping.html).  
  
***Current pexpo's version: 1.30***  
(scroling host-list!)

## Install
Install Page: https://github.com/nnao45/pexpo/releases/tag/1.30

## Usage
```bash
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
	(If you want to Request, "http and https", Using Both -H & -S.)
	   
	-H This optison is like "curl". So you Sending HTTP(:80) GET Request instead of the PING...!
	   
	-S This optison is like "curl". So you Sending HTTP"S"(:443) GET Request instead of the PING...!
	
	-H or -S options is sending HTTP/HTTPS GET Request instead of the PING.
	(Just like, curl -LIs www.google.com -o /dev/null -w '%{http_code}\n')
	This Request is ververy simple GET Request, Only Getting status code(No header, No form, No getting data.)

	And, if http status code is "200", string color is Blue, else Red.
```
 
## Demo (macOS 10.13.1):apple:
![result](https://github.com/nnao45/naoGifRepo/blob/master/pexpo-mac.gif)
 
## Demo (Ubuntu16.04):penguin:
![result](https://github.com/nnao45/naoGifRepo/blob/master/pexpo_1.20_linux.gif)

## Demo (windows10):four_leaf_clover:
![result](https://github.com/nnao45/naoGifRepo/blob/master/pexpo_1.20_win.gif)

## Demo (windows10 & HTTPING):earth_asia:
![result](https://github.com/nnao45/naoGifRepo/blob/master/pexpo_HS_1.20_wins.gif)

## Implementation
- Very light, and quick application(for Sending ICMP to the too many hosts):metal:
- ONLY one app run on multi platforms(Windows7, Windows10, Mac, Linux...)!!:kissing_heart:
- You can send ICMP or HTTP GET or HTTPS GET ipv4, and ipv6!!:open_mouth:
- pexpo has several options. You can change ping interval, timeout, select ping-list, ,help Cisco using, & http ping mode!:octocat:
- Display Counting Ping loss per host:point_up_2:
- Display Current Dead host(if host is revive, and dead mark will be vanish):boom:
- pexpo has Pausing Implementation. if you want, push "Crtl+S":traffic_light:
### more...
- Logging ping result($HOME/.pexpo/result_$DATE_.txt).
- Check the syntax on the ping-list(# is comment out, ignoring blank line, using tab is ok, no description is ok...).
- Push ArrowUp or Down key, scroll host-list :arrow_double_up: :arrow_double_down:
  
## Release note
- version 1.30...Scroling host-list!!!!! :fish_cake:
- version 1.25...travis support & reading ping-list's bug fix.
- version 1.24...go report A+!!(no implement change)
- version 1.23...Little performance up(assign cap in the string[])
- version 1.22...Too Little bug fix(string join -> append []string)
- version 1.21...Too Little change in code & icon+
- version 1.20...Wow!!:heart_eyes:Adding "HTTP PING"!!
  - "-H", http_ping "-S", https_ping. Using Both, Sending Both with not error.
  - Accompanied by the http ping implementation, little change variavle, channel. There is no change in ICMP behavior.
- version 1.10...Too little additinal change under line, There is no change in Basic behavior.
  - Print version.
  - Readability up(not using global variable).
  - Add channel, "received"(Both directions key interrupt channels).
  - When push Ctrl+S, change key interrupt message.
- version 1.00...Implementated Basic functions.
  
***Have a nice go hacking days***:sparkles::wink:
## Writer & License
pexpo was writed by nnao45 (WORK:Network Engineer, Twitter:@A_Resas, MAIL:n4sekai5y@gmail.com).  
This software is released under the MIT License, see LICENSE.
