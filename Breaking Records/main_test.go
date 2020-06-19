package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		scores []int
		max int
		min int
	}{
		{
			scores: []int{10, 5, 20, 20, 4, 5, 2, 25, 1},
			max: 2,
			min: 4,
		},
	}

	for _, test := range tests {
		if max, min := breakingRecords(test.scores); max != test.max || min != test.min {
			t.Errorf(
				"Given Scores %v\n; Expected %v %v; Got %v %v",
				test.scores, test.max, test.min, max, min,
			)
		}
	}
}
