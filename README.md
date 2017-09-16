# pexpo
pexpo is ping sending tool with cool clomun & logging loss count int the Multi platform(Windos, Mac, Linux...).  
Sendig TUI Engine is ![termbox-go](https://github.com/nsf/termbox-go), Sending ICMP Engine is ![go-fastping](https://github.com/tatsushid/go-fastping)
This is inspired [pinger](https://github.com/hirose31/pinger), [Exping](http://www.woodybells.com/exping.html).
  
## Demo (Linux)
![result](https://github.com/nao4arale/pexpo/blob/master/pexpo_linux.gif)

## Demo (windows)
![result](https://github.com/nao4arale/pexpo/blob/master/pexpo_windows.gif)

## implemention
- very light application.
- supporting the windows!!
- ipv4, and ipv6!!
- logging ping result($HOME/.pexpo/result_$DATE_.txt)
- Display Counting Ping loss per host.
- Display Current Dead host(if host is revive, And dead mark is vanish).

## Writer & License
pexpo's was writed by nao4arale (Twitter:@A_Resas, MAIL:n4sekai5y@gmail.com).  
This software is released under the MIT License, see LICENSE.
