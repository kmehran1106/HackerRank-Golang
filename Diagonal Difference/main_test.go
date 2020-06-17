package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		output int
	}{
		{matrix: [][]int{
			{11, 2, 4},
			{4, 5, 6},
			{10, 8, -12},
		}, output: 15},
	}

	for _, test := range tests {
		if got := diagonalDifference(test.matrix); got != test.output {
			t.Errorf(
				"For Matrix %v; Expected Output %v; Got %v",
				test.matrix, test.output, got,
			)
		}
	}
}
