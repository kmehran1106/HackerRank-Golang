package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}},
			output: 7,
		},
	}

	for _, test := range tests {
		if got := magicSquareForming(test.input); got != test.output {
			t.Errorf("magicSquareForming Got%v while Expecting %v", got, test.output)
		}
	}
}
