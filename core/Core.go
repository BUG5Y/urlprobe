package core

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
	"strings"
	"golang.org/x/sync/semaphore"
	"golang.org/x/net/context"
)

type Config struct {
	MaxConcurrency int
	Timeout        time.Duration
	Protocols      []string
	Ports          []string
}

var (
	MaxConcurrency = 20
	Timeout = 15*time.Second
	Protocols = []string{"http", "https"}
	Ports = []string{"80", "443", "8080", "8443", "3000", "4443", "8000", "8888", "5000", "8008"}
)

func Start(filePath string) {
	urls, err := readURLs(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	config := setConfiguration()

	sem := semaphore.NewWeighted(int64(config.MaxConcurrency))
	ctx := context.Background()

	var wg sync.WaitGroup
	results := make(chan string, len(urls)*len(config.Protocols)*len(config.Ports))

    for _, url := range urls {
        host := strings.Split(url, ":")[0] // Extract host from URL
        ip := resolveIP(host)
        
        for _, protocol := range config.Protocols {
            for _, port := range config.Ports {
                wg.Add(1)
                go func(u, p string, port string, ip string) {
                    defer wg.Done()
                    if err := sem.Acquire(ctx, 1); err != nil {
                        fmt.Printf("Failed to acquire semaphore: %v\n", err)
                        return
                    }
                    defer sem.Release(1)

                    active := sendRequest(p, u, port, config.Timeout)
                    if active {
                        results <- fmt.Sprintf("\033[32m[Active URL] %s://%s:%s - IP: %s\033[0m", p, u, port, ip)
                    } else {
                        results <- fmt.Sprintf("[No Response] %s://%s:%s - IP: %s", p, u, port, ip)
                    }
                }(url, protocol, port, ip)
            }
        }
    }

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func readURLs(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls, scanner.Err()
}

func setConfiguration() Config {
	config := Config{
		MaxConcurrency: MaxConcurrency,
		Timeout:        Timeout,
		Protocols:      Protocols,
		Ports:          Ports,
	}
	return config
}