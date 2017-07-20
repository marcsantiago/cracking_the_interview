package main

import (
	"bytes"
	"strconv"
)

// global buf incase the compress function is called many times
// this will save memory allocations
var buf bytes.Buffer

func writeWrapper(buf *bytes.Buffer, r rune, counter int) {
	sCounter := strconv.Itoa(counter)
	buf.WriteRune(r)
	buf.WriteString(sCounter)
	return
}

func compress(s string) string {
	defer buf.Reset()
	counter := 1
	for i, r := range s {
		f := i + 1
		if f < len(s) {
			if r == rune(s[f]) {
				counter++
			} else {
				writeWrapper(&buf, r, counter)
				counter = 1
			}
		} else {
			writeWrapper(&buf, r, counter)
			counter = 1
		}
	}
	if buf.Len() < len(s) {
		return buf.String()
	}
	return s
}

func main() {}
