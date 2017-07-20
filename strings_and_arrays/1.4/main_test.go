package main

import "testing"

var cases = []struct {
	Input  string
	Output string
}{
	{"mr john smith    ", "mr%20john%20smith"},
	{"mr john  ", "mr%20john"},
	{"mr john smith fell in a cup            ", "mr%20john%20smith%20fell%20in%20a%20cup"},
}

func TestInplaceSwap(t *testing.T) {
	for _, c := range cases {
		if len(c.Input) != len(c.Output) {
			t.Fatalf("check test cases except length of %d got %d", len(c.Output), len(c.Input))
		}
		out := inplaceSwap([]rune(c.Input))
		if out != c.Output {
			t.Errorf("Expected: %s\ngot: %s\n", c.Output, out)
		}
	}
}
