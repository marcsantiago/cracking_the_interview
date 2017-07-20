package main

import "testing"

var cases = []struct {
	In  string
	Out string
}{
	{"aabcccccaaa", "a2b1c5a3"},
}

func TestCompression(t *testing.T) {
	for _, c := range cases {
		out := compress(c.In)
		if out != c.Out {
			t.Errorf("got: %s expected: %s\n", out, c.Out)
		}
	}
}
