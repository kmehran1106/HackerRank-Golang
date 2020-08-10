package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		array  []int
		output int
	}{
		{
			array:  []int{2, 3, 4, 5, 6},
			output: 4,
		},
		{
			array:  []int{1, 2},
			output: 0,
		},
	}

	for _, test := range tests {
		if got := fairRations(test.array); got != test.output {
			t.Errorf(
				"For array=%v; Got %v while Expecting %v",
				test.array, got, test.output,
			)
		}
	}
}
