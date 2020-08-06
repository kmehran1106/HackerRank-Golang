package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		n      int
		k      int
		arr    []int
		output int
	}{
		{
			n:      5,
			k:      3,
			arr:    []int{4, 2, 6, 1, 10},
			output: 4,
		},
		{
			n:      10,
			k:      5,
			arr:    []int{3, 8, 15, 11, 14, 1, 9, 2, 24, 31},
			output: 8,
		},
		{
			n: 40,
			k: 7,
			arr: []int{1, 10, 12, 4, 11, 6, 8, 15, 23, 24, 23, 24, 39, 34, 50, 3,
				58, 62, 71, 79, 95, 100, 2, 2, 100, 100, 100, 100, 100, 100, 1,
				100, 100, 100, 100, 100, 3, 100, 100, 100,
			},
			output: 12,
		},
	}

	for _, test := range tests {
		if got := getWorkbookSpecialProblems(test.n, test.k, test.arr); got != test.output {
			t.Errorf(
				"For n=%v k=%v arr=%v; Got %v while Expecting %v",
				test.n, test.k, test.arr, got, test.output,
			)
		}
	}
}
