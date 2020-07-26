package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		d      int
		array  []int
		output int
	}{
		{
			d:      3,
			array:  []int{1, 2, 4, 5, 7, 8, 10},
			output: 3,
		},
		{
			d:      3,
			array:  []int{1, 6, 7, 7, 8, 10, 12, 13, 14, 19},
			output: 2,
		},
	}

	for _, test := range tests {
		if got := beautifulTriplets(test.d, test.array); got != test.output {
			t.Errorf(
				"Given Difference %v and Array %v; Expected %v; Got %v",
				test.d, test.array, test.output, got,
			)
		}
	}
}
