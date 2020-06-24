package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  []string
		output int
	}{
		{input: []string{"U", "D", "D", "D", "U", "D", "U", "U"}, output: 1},
	}

	for _, test := range tests {
		if got := countingValleys(test.input); got != test.output {
			t.Errorf(
				"For Input %v; Expected Output %v; Got %v",
				test.input, test.output, got,
			)
		}
	}
}
