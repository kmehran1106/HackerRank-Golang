package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{2, 3, 1},
			output: []int{2, 3, 1},
		},
		{
			input:  []int{4, 3, 5, 1, 2},
			output: []int{1, 3, 5, 4, 2},
		},
	}

	for _, test := range tests {
		got := sequenceEquation(test.input)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf("sequenceEquation Got%v while Expecting %v", got, test.output)
			}
		}
	}
}
