package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		grid   []string
		output []string
	}{
		{
			grid:   []string{"1112", "1912", "1892", "1234"},
			output: []string{"1112", "1X12", "18X2", "1234"},
		},
		{
			grid:   []string{"12", "12"},
			output: []string{"12", "12"},
		},
	}

	for _, test := range tests {
		got := cavityMap(test.grid)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf(
					"For grid=%v; Got %v while expecting %v",
					test.grid, got, test.output,
				)
			}
		}
	}
}
