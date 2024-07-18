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
[No Response] http://youtube.com:443 - IP: 142.250.72.78
[No Response] https://yahoo.com:80 - IP: 74.6.231.20
[No Response] https://youtube.com:80 - IP: 142.250.72.78
[Active URL] http://youtube.com:80 - IP: 142.250.72.78
[Active URL] https://youtube.com:443 - IP: 142.250.72.46
[Active URL] https://yahoo.com:443 - IP: 69.147.65.252
[No Response] https://youtube.com:3000 - IP: 
[No Response] http://youtube.com:8080 - IP: 
[No Response] https://youtube.com:4443 - IP:

```
