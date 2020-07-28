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
			input:  []int{7, 1, 3, 4, 1, 7},
			output: 3,
		},
		{
			input:  []int{1, 2, 3, 4, 10},
			output: -1,
		},
	}

	for _, test := range tests {
		if got := minimumDistances(test.input); got != test.output {
			t.Errorf(
				"For Input %v; Got %v while Expecting %v",
				test.input, got, test.output,
			)
		}
	}
}
