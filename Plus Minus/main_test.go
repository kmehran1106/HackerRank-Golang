package main

import (
	"testing"
)


func TestCode(t *testing.T) {
	var tests = []struct {
		input []int
		output []float64
	} {
		{
			input: []int{1, 0, -1},
			output: []float64{1.0/3.0, 1.0/3.0, 1.0/3.0},
		},
		{
			input: []int{1, 1, 1, 0, 0, -1, -1},
			output: []float64{3.0/7.0, 2.0/7.0, 2.0/7.0},
		},
	}

	for _, test := range tests {
		got := PlusMinus(test.input)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf("PlusMinus Got%v while Expecting %v", got, test.output)
			}
		}
	}
}