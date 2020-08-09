package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		n      int
		arr    []int
		output int
	}{
		{
			n:      5,
			arr:    []int{0, 4},
			output: 2,
		},
		{
			n:      6,
			arr:    []int{0, 1, 2, 3, 4, 5},
			output: 0,
		},
	}

	for _, test := range tests {
		if got := getSpaceStationMax(test.n, test.arr); got != test.output {
			t.Errorf(
				"For n=%v arr=%v; Got %v while Expecting %v",
				test.n, test.arr, got, test.output,
			)
		}
	}
}
