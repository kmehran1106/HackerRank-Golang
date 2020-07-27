package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  []string
		output []int
	}{
		{
			input:  []string{"10101", "11100", "11010", "00101"},
			output: []int{5, 2},
		},
		{
			input:  []string{"11101", "10101", "11001", "10111", "10000", "01110"},
			output: []int{5, 6},
		},
	}

	for _, test := range tests {
		got := getACMTeam(test.input)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf(
					"For Input %v; Expected Output %v; Got %v",
					test.input, test.output, got,
				)
			}
		}
	}
}
