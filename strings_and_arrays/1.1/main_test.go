package main

import "testing"

var cases = []struct {
	In       string
	Excepted bool
}{
	{"", true},
	{"h", true},
	{"hh", false},
	{"heh", false},
	{"hello", false},
	{"marc", true},
	{"Candy", true},
	{"canada", false},
	{"Hh", false},
	{"heE", false},
}

func TestWithMap(t *testing.T) {
	for _, c := range cases {
		got := isUniqueWithHashMap(c.In)
		if got != c.Excepted {
			t.Errorf("for the entered string %s the expected value was %v got %v\n", c.In, c.Excepted, got)
		}
	}
}

func TestWithOutMap(t *testing.T) {
	for _, c := range cases {
		got := isUniqueWithOutHashMap(c.In)
		if got != c.Excepted {
			t.Errorf("for the entered string %s the expected value was %v got %v\n", c.In, c.Excepted, got)
		}
	}
}
