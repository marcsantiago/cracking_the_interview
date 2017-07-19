package main

import "testing"

var cases = []struct {
	A        string
	B        string
	Expected bool
}{
	{"bat", "tab", true},
	{"a", "b", false},
	{"shit", "this", true},
	{"can", "bottle", false},
}

func TestPerm(t *testing.T) {
	for _, c := range cases {
		got := isPermutation(c.A, c.B)
		if c.Expected != got {
			t.Errorf("For string %s and %s expected %v got %v\n", c.A, c.B, c.Expected, got)
		}
	}
}
