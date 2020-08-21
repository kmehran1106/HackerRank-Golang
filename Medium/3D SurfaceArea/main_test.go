package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		grid   [][]int
		output int
	}{
		{
			grid: [][]int{
				{1},
			},
			output: 6,
		},
		{
			grid: [][]int{
				{1, 3, 4},
				{2, 2, 3},
				{1, 2, 4},
			},
			output: 60,
		},
	}

	for _, test := range tests {
		if got := surfaceArea(test.grid); got != test.output {
			t.Errorf("Surface Area Function Got%v while Expecting %v", got, test.output)
		}
	}
}
