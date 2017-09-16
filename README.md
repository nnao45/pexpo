# pexpo
pexpo is ping sending tool with cool clomun & logging loss count int the Multi platform(Windows, Mac, Linux...).  
Sendig TUI Engine is ![termbox-go](https://github.com/nsf/termbox-go), Sending ICMP Engine is ![go-fastping](https://github.com/tatsushid/go-fastping).  
So, pexpo's code is NATIVE ![The Go Programming Language](http://golang.org) application.  
This is inspired the [pinger](https://github.com/hirose31/pinger), [Exping](http://www.woodybells.com/exping.html).
  
## Demo (Linux)
![result](https://github.com/nao4arale/pexpo/blob/master/pexpo_linux.gif)

## Demo (windows10)
![result](https://github.com/nao4arale/pexpo/blob/master/pexpo_windows.gif)

## implemention
- very light application(for Pinging to the too many hosts).
- Run on multi platforms(Windows7, Windows10, Mac, Linux...) with ONLY one app!!
- ipv4, and ipv6!!
- logging ping result($HOME/.pexpo/result_$DATE_.txt).
- Display Counting Ping loss per host.
- Display Current Dead host(if host is revive, and dead mark is vanish).

## Writer & License
pexpo's was writed by nao4arale (Twitter:@A_Resas, MAIL:n4sekai5y@gmail.com).  
This software is released under the MIT License, see LICENSE.
