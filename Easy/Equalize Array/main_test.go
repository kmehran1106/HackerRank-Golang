package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{
		{input: []int{3, 3, 2, 1, 3}, output: 2},
		{input: []int{1, 2, 3, 1, 2, 3, 3, 3}, output: 4},
	}

	for _, test := range tests {
		if got := equalizeArray(test.input); got != test.output {
			t.Errorf(
				"For Input %v; Expected Rounded Input %v; Got %v",
				test.input, test.output, got,
			)
		}
	}
}
