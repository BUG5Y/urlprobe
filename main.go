package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/bug5y/urlprobe/core"
)

func main() {
	filePath := flag.String("f", "", "Path to the input file")
	threads := flag.Int("t", 20, "Max threads")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Error: File path is required. Use -f flag to specify the file path.")
		flag.Usage()
		os.Exit(1)
	}

	core.Start(*filePath, *threads)
}
