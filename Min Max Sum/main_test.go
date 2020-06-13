package main

import (
	"testing"
)


func TestCode(t *testing.T) {
	var tests = []struct {
		input []int64
		output []int64
	} {
		{
			input: []int64{1, 2, 3, 4, 5},
			output: []int64{10, 14},
		},
	}

	for _, test := range tests {
		got := MinMaxSum(test.input)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf("PlusMinus Got%v while Expecting %v", got, test.output)
			}
		}
	}
}