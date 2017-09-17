# pexpo
pexpo is ping sending tui tool with cool clomun & logging loss-count in the multi platforms(Windows, Mac, Linux...).  
pexpo has tui engine is [termbox-go](https://github.com/nsf/termbox-go), sending ICMP engine is [go-fastping](https://github.com/tatsushid/go-fastping).  
So, pexpo's code is NATIVE [The Go Programming Language](http://golang.org) application. 
This is inspired the [pinger](https://github.com/hirose31/pinger), [Exping](http://www.woodybells.com/exping.html).
  
## Usage
```bash
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
```
  
## Demo (Linux)
![result](https://github.com/nao4arale/pexpo/blob/master/pexpo_linux.gif)

## Demo (windows10)
![result](https://github.com/nao4arale/pexpo/blob/master/pexpo_win.gif)

## Implementation
- Very light, and quick application(for Sending ICMP to the too many hosts):metal:
- ONLY one app run on multi platforms(Windows7, Windows10, Mac, Linux...)!!:kissing_heart:
- You can send ICMP ipv4, and ipv6!!:open_mouth:
- pexpo has several options. You can hange ping interval, timeout, selecting ping-list, & help Cisco using:octocat:
- Display Counting Ping loss per host:point_up_2:
- Display Current Dead host(if host is revive, and dead mark is vanish):boom:
### more...
- logging ping result($HOME/.pexpo/result_$DATE_.txt).
- Check the syntax on the ping-list(# is comment out, ignoring blank line, no description is ok...).
  
Have a nice go hacking days:sparkles::wink:
## Writer & License
pexpo's was writed by nao4arale (Twitter:@A_Resas, MAIL:n4sekai5y@gmail.com).  
This software is released under the MIT License, see LICENSE.
