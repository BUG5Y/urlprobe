package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/bug5y/urlprobe/core"
)

func main() {
	var config core.Config
	defaultTimeout := 15 * time.Second
	defaultThreads := 20
	defaultProtocols := []string{"http", "https"}
	defaultPorts := []string{"80", "443", "8080", "8443", "3000", "4443", "8000", "8888", "5000", "8008"}
	defaultUserAgent := "Mozilla/5.0 (compatible; urlprobe/1.0)"

	filePath := flag.String("f", "", "Path to the input file")
	threads := flag.Int("t", defaultThreads, "Max threads")
	flag.DurationVar(&config.Timeout, "timeout", defaultTimeout, "Timeout duration")
	protocols := flag.String("p", strings.Join(defaultProtocols, ","), "Comma-separated list of protocols")
	ports := flag.String("ports", strings.Join(defaultPorts, ","), "Comma-separated list of ports")
	userAgent := flag.String("header", defaultUserAgent, "Custom user agent to add to each request")
	help := flag.Bool("h", false, "Show help")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *filePath == "" {
		fmt.Println("Error: File path is required. Use -f flag to specify the file path.")
		flag.Usage()
		os.Exit(1)
	} else {
		config.FilePath = *filePath
	}

	if *userAgent != "" {
		config.UserAgent = *userAgent
	}

	config.MaxConcurrency = *threads
	config.Protocols = strings.Split(*protocols, ",")
	config.Ports = strings.Split(*ports, ",")

	core.Start(config)
}
