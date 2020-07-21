package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  [][]int
		output bool
	}{
		{
			input: [][]int{
				{0, 2, 1},
				{1, 1, 1},
				{2, 0, 0},
			},
			output: true,
		},
		{
			input: [][]int{
				{1, 3, 1},
				{2, 1, 2},
				{3, 3, 3},
			},
			output: false,
		},
	}

	for _, test := range tests {
		if got := organizingContainers(test.input); got != test.output {
			t.Errorf(
				"Organzing Containers for Input=%v; Got%v; Expected %v",
				test.input, got, test.output,
			)
		}
	}

}
