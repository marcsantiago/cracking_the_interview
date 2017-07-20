package main

import "testing"

var cases = []struct {
	Input      string
	Output     string
	TrueLength int
}{
	{"mr john smith    ", "mr%20john%20smith", 13},
	{"mr john  ", "mr%20john", 7},
	{"mr john smith fell in a cup            ", "mr%20john%20smith%20fell%20in%20a%20cup", 27},
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
