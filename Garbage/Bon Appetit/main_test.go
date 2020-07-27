package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		array  []int
		index  int
		paid   int
		output int
	}{
		{[]int{3, 10, 2, 9}, 1, 12, 5},
		{[]int{3, 10, 2, 9}, 1, 7, 0},
	}

	for _, test := range tests {
		if got := bonAppetit(test.array, test.index, test.paid); got != test.output {
			t.Errorf(
				"Given Array %v, Index %v, Paid %v\n; Expected %v; Got %v",
				test.array, test.index, test.paid, test.output, got,
			)
		}
	}
}
