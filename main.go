package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// flag to count lines instead
	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("b", false, "count bytes")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	// by default scanner splits lines
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	// count count byte
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	c := 0

	for scanner.Scan() {
		c++
	}

	return c
}
