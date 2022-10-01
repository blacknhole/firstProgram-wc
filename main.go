package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	files := flag.Bool("f", false, "Read data from files")
	flag.Parse()

	if *files {
		for i, filename := range flag.Args() {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Printf("%d. %s: %d\n", i+1, filename, count(file, *lines, *bytes))
		}
	} else {
		fmt.Println(count(os.Stdin, *lines, *bytes))
	}
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	wc := 0
	scanner := bufio.NewScanner(r)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}
	for scanner.Scan() {
		wc++
	}
	return wc
}
