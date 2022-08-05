package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	wc := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wc++
	}
	return wc
}
