package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		threshold int
		timings   []int
		output    bool
	}{
		{
			threshold: 3,
			timings:   []int{-1, -3, 4, 2},
			output:    true,
		},
		{
			threshold: 2,
			timings:   []int{0, -1, 2, 1},
			output:    false,
		},
	}

	for _, test := range tests {
		if got := angryProfessorCancelsClass(test.threshold, test.timings); got != test.output {
			t.Errorf(
				"Given Threshold %v and Timings %v\n; Expected %v; Got %v",
				test.threshold, test.timings, test.output, got,
			)
		}
	}
}
