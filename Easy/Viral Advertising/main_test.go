package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  int
		output int
	}{
		{input: 5, output: 24},
		{input: 3, output: 9},
	}

	for _, test := range tests {
		if got := viralAdvertising(test.input); got != test.output {
			t.Errorf(
				"Given Input %v; Expected %v; Got %v",
				test.input, test.output, got,
			)
		}
	}
}
