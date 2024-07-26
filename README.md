# urlprobe

Probe a list of domains for working web servers on common ports.

## Install

```
▶ go install github.com/bug5y/urlprobe@latest
```

## Example Usage

```
▶ urlprobe -h
Usage of urlprobe:
  -f string
        Path to the input file
  -h    Show help
  -header string
        Custom user agent to add to each request. Default: Mozilla/5.0 (Windows NT 10.0; Android 13; Mobile; rv:120.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 OPR/95.0.0.0 (default "Mozilla/5.0 (Windows NT 10.0; Android 13; Mobile; rv:120.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 OPR/95.0.0.0")
  -p string
        Comma-separated list of protocols (default "http,https")
  -ports string
        Comma-separated list of ports (default "80,443,8080,8443,3000,4443,8000,8888,5000,8008")
  -t int
        Max threads (default 20)
  -timeout duration
        Timeout duration (default 15s)

```

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
[No Response] http://yahoo.com:443 - IP: 98.137.11.163
```
