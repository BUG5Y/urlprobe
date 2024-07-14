package main

import (
	"flag"
	"fmt"
	"os"
	"urlprobe/core"
)

func main() {
	filePath := flag.String("f", "", "Path to the input file")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Error: File path is required. Use -f flag to specify the file path.")
		flag.Usage()
		os.Exit(1)
	}

	core.Start(*filePath)
}