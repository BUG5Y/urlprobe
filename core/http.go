package core

import (
	"time"
	"net/http"
	"fmt"
	"net"
	"strings"
	"strconv"
)

func resolveIP(host string) string {
    ips, err := net.LookupIP(host)
    if err != nil || len(ips) == 0 {
        return "IP not found"
    }
    var ipStrings []string
    for _, ip := range ips {
        ipStrings = append(ipStrings, ip.String())
    }
    return strings.Join(ipStrings, ", ")
}

func sendRequest(protocol, url string, portStr string, timeout time.Duration) bool {
	port, err := strconv.Atoi(portStr)
    if err != nil {
        fmt.Println("Error:", err)
        return false
    }

    client := &http.Client{
        Timeout: timeout,
    }

    fullURL := fmt.Sprintf("%s://%s:%d", protocol, url, port)
    
    req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
        return false
    }

    resp, err := client.Do(req)
    if err != nil {
        return false
    }
    defer resp.Body.Close()

    return resp.StatusCode >= 200 && resp.StatusCode < 500
}


