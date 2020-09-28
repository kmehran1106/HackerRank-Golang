package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		seconds int
		grid    []string
		output  []string
	}{
		{
			seconds: 3,
			grid: []string{
				".......",
				"...O...",
				"....O..",
				".......",
				"OO.....",
				"OO.....",
			},
			output: []string{
				"OOO.OOO",
				"OO...OO",
				"OOO...O",
				"..OO.OO",
				"...OOOO",
				"...OOOO",
			},
		},
		{
			seconds: 5,
			grid: []string{
				".......",
				"...O.O.",
				"....O..",
				"..O....",
				"OO...OO",
				"OO.O...",
			},
			output: []string{
				".......",
				"...O.O.",
				"...OO..",
				"..OOOO.",
				"OOOOOOO",
				"OOOOOOO",
			},
		},
	}

	for _, test := range tests {
		got := bomberman(test.seconds, test.grid)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf("Bomberman Function Got%v while Expecting %v", got, test.output)
			}
		}
	}
}
