package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		heights []int
		input   string
		output  int
	}{
		{
			heights: []int{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
			input:   "abc",
			output:  9,
		},
	}

	for _, test := range tests {
		if got := pdfViewer(test.heights, test.input); got != test.output {
			t.Errorf(
				"Given Heights %v\nInput %v; Expected Output %v; Got %v",
				test.heights, test.input, test.output, got,
			)
		}
	}
}
