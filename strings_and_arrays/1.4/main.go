package main

import (
	"strings"
)

const space = rune(' ')

func easySwap(s string) string {
	return strings.Replace(strings.TrimSpace(s), " ", "%20", -1)
}

// below will still fail if more than one space is used between words
// also I assume that the number of spaces at the end is equal to the
// exact spaces of numbers needed
func getStartEnd(s []rune) (start, end int) {
	for i, r := range s {
		f := i + 1
		if f < len(s) {
			if r == space && s[f] != space {
				start = f
			}
			if r == space && s[f] == space {
				end = i
			}
			if start != 0 && end != 0 {
				break
			}
		}
	}
	return
}

func countSpacesAfterEnd(s []rune, end int) (spaces int) {
	for _, r := range s[end:] {
		if r == space {
			spaces++
		}
	}
	return
}

func replace(s []rune) {
	fill := []rune{'0', '2', '%'}
	var counter int
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == space {
			s[i] = fill[counter]
			counter++
		}
		if counter == 3 {
			return
		}
	}
}

func inplaceSwap(s []rune) string {
	for {
		start, end := getStartEnd(s)
		if start == 0 && end == 0 {
			break
		}
		numberOfSpaces := countSpacesAfterEnd(s, end)
		for i := end - 1; i+1 > start; i-- {
			r := s[i]
			newPos := i + numberOfSpaces
			s[newPos] = r
			s[i] = space
		}
		replace(s)
	}
	return string(s)
}

func main() {

}
