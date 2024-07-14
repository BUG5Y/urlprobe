# urlprobe

Probe a list of domains for working web servers on common ports.

## Install

```
▶ go install github.com/BUG5Y/urlprobe@latest
```

## Usage

urlprobe accepts the file path through the -f flag:

```
▶ cat example.txt                                                                           
youtube.com
facebook.com
yahoo.com
▶ urlprobe -f example.txt
[No Response] http://youtube.com:443 - IP: 142.250.191.238, 2607:f8b0:4009:80b::200e
[No Response] https://youtube.com:80 - IP: 142.250.191.238, 2607:f8b0:4009:80b::200e
[Active URL] http://youtube.com:80 - IP: 142.250.191.238, 2607:f8b0:4009:80b::200e
[Active URL] https://youtube.com:443 - IP: 142.250.191.238, 2607:f8b0:4009:80b::200e
[No Response] http://facebook.com:443 - IP: 157.240.254.35, 2a03:2880:f175:181:face:b00c:0:25de
[Active URL] http://facebook.com:80 - IP: 157.240.254.35, 2a03:2880:f175:181:face:b00c:0:25de
[No Response] https://youtube.com:8080 - IP: 142.250.191.238, 2607:f8b0:4009:80b::200e
[No Response] https://youtube.com:4443 - IP: 142.250.191.238, 2607:f8b0:4009:80b::200e
[No Response] http://youtube.com:4443 - IP: 142.250.191.238, 2607:f8b0:4009:80b::200e
[No Response] http://youtube.com:8080 - IP: 142.250.191.238, 2607:f8b0:4009:80b::200e

```
