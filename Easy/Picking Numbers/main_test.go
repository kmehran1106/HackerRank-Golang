package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 2, 1, 2},
			output: 5,
		},
		{
			input:  []int{66, 66, 66, 66, 66, 66},
			output: 6,
		},
	}

	for _, test := range tests {
		if got := pickingNumbers(test.input); got != test.output {
			t.Errorf(
				"Input Array %v; Expected %v; Got %v",
				test.input, test.output, got,
			)
		}
	}
}
