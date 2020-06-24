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
			input:  []int{1, 1, 2, 2, 3},
			output: 1,
		},
		{
			input:  []int{1, 4, 4, 4, 5, 3},
			output: 4,
		},
	}

	for _, test := range tests {
		if got := migratoryBirds(test.input); got != test.output {
			t.Errorf(
				"Input Array of Birds %v; Expected %v; Got %v",
				test.input, test.output, got,
			)
		}
	}
}
