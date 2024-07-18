package core

import (
	"time"
	"net/http"
	"fmt"
	"net"
	"strconv"
    "context"
)

func sendRequest(protocol, url string, portStr string, timeout time.Duration) (bool, string) {
    port, err := strconv.Atoi(portStr)
    if err != nil {
        fmt.Println("Error:", err)
        return false, ""
    }

    var ipAddr string
    dialer := &net.Dialer{
        Timeout: timeout,
    }

    transport := &http.Transport{
        DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
            conn, err := dialer.DialContext(ctx, network, addr)
            if err != nil {
                return nil, err
            }
            if tcpConn, ok := conn.(*net.TCPConn); ok {
                if remoteAddr, ok := tcpConn.RemoteAddr().(*net.TCPAddr); ok {
                    ipAddr = remoteAddr.IP.String()
                }
            }
            return conn, nil
        },
    }

    client := &http.Client{
        Timeout:   timeout,
        Transport: transport,
    }

    fullURL := fmt.Sprintf("%s://%s:%d", protocol, url, port)

    req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
        return false, ipAddr
    }

    resp, err := client.Do(req)
    if err != nil {
        return false, ipAddr
    }
    defer resp.Body.Close()

    return resp.StatusCode >= 200 && resp.StatusCode < 500, ipAddr
}
