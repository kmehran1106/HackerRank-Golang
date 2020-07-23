package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"ab", "ba"},
		{"bb", ""},
		{"hefg", "hegf"},
		{"dhck", "dhkc"},
		{"dkhc", "hcdk"},
		{"lmno", "lmon"},
		{"dcba", ""},
		{"dcbb", ""},
		{"abdc", "acbd"},
		{"abcd", "abdc"},
		{"fedcbabcd", "fedcbabdc"},
	}

	for _, test := range tests {
		if got := biggerIsGreater(test.input); got != test.output {
			t.Errorf(
				"Bigger is Greater for Input=%v; Got%v; Expected %v",
				test.input, got, test.output,
			)
		}
	}

}
