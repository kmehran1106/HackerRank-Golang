package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		grid  []string
		pattern []string
		output bool
	}{
		{
			grid: []string{
				"7283455864",
				"6731158619",
				"8988242643",
				"3830589324",
				"2229505813",
				"5633845374",
				"6473530293",
				"7053106601",
				"0834282956",
				"4607924137",
			},
			pattern: []string{
				"9505",
				"3845",
				"3530",
			},
			output: true,
		},
	}

	for _, test := range tests {
		if got := gridSearch(test.grid, test.pattern); got != test.output {
			t.Errorf("GridSearch Function Got%v while Expecting %v", got, test.output)
		}
	}
}
