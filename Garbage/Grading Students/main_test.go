package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  int
		output int
	}{
		{input: 74, output: 75},
		{input: 33, output: 33},
		{input: 76, output: 76},
	}

	for _, test := range tests {
		if got := gradingStudents(test.input); got != test.output {
			t.Errorf(
				"For Input %v; Expected Rounded Input %v; Got %v",
				test.input, test.output, got,
			)
		}
	}
}
