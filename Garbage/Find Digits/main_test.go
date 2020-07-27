package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  int
		output int
	}{
		{input: 114108089, output: 3},
		{input: 110110015, output: 6},
		{input: 121, output: 2},
	}

	for _, test := range tests {
		if got := findDigits(test.input); got != test.output {
			t.Errorf(
				"For Input %v; Expected Rounded Input %v; Got %v",
				test.input, test.output, got,
			)
		}
	}
}
