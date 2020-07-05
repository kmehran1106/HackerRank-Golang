package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{5, 4, 4, 2, 2, 8},
			output: []int{6, 4, 2, 1},
		},
		{
			input:  []int{1, 2, 3, 4, 3, 3, 2, 1},
			output: []int{8, 6, 4, 1},
		},
	}

	for _, test := range tests {
		got := cutSticks(test.input)
		for i, v := range got {
			if test.output[i] != v {
				t.Errorf(
					"For Input=%v; Expected Output=%v; Got %v",
					test.input, test.output, got,
				)
			}
		}
	}
}
