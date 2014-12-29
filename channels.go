package helpers

import (
	"bufio"
	"fmt"
	"io"
)

// NewWriterChannel returns a send only channel, that writes everything it
// receives into w using fmt.Fprintln and adding newlines
func NewWriterChannel(w io.Writer) chan<- string {
	wc := make(chan string)

	go func() {
		for signal := range wc {
			fmt.Fprintln(w, signal)
		}
	}()

	return wc
}

// NewScannerChannel returns a receive only channel, that reads lines from
// r using a bufio.Scanner
func NewScannerChannel(r io.Reader) <-chan string {
	rc := make(chan string)

	go func() {
		s := bufio.NewScanner(r)

		// Send every line
		for s.Scan() {
			rc <- s.Text()
		}

		close(rc)
	}()

	return rc
}
