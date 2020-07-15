package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		n         int
		k         int
		rQ        int
		cQ        int
		obstacles [][]int
		output    int
	}{
		{
			4, 0, 4, 4, [][]int{}, 9,
		},
		{
			5, 3, 4, 3, [][]int{{5, 5}, {4, 2}, {2, 3}}, 10,
		},
		{
			1, 0, 1, 1, [][]int{}, 0,
		},
	}

	for _, test := range tests {
		if got := getQueensAttack(test.n, test.k, test.rQ, test.cQ, test.obstacles); got != test.output {
			t.Errorf(
				"Queens Attack II for n=%v, k=%v, r_q=%v, c_q=%v, obstacles=%v; Got%v; Expected %v",
				test.n, test.k, test.rQ, test.cQ, test.obstacles, got, test.output,
			)
		}
	}

}
