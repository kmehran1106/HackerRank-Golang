package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		widthArray []int
		lowerIndex int
		upperIndex int
		output     int
	}{
		{
			widthArray: []int{2, 3, 1, 2, 3, 2, 3, 3},
			lowerIndex: 0,
			upperIndex: 3,
			output:     1,
		},
		{
			widthArray: []int{2, 3, 1, 2, 3, 2, 3, 3},
			lowerIndex: 4,
			upperIndex: 6,
			output:     2,
		},
		{
			widthArray: []int{2, 3, 1, 2, 3, 2, 3, 3},
			lowerIndex: 6,
			upperIndex: 7,
			output:     3,
		},
		{
			widthArray: []int{2, 3, 1, 2, 3, 2, 3, 3},
			lowerIndex: 3,
			upperIndex: 5,
			output:     2,
		},
		{
			widthArray: []int{2, 3, 1, 2, 3, 2, 3, 3},
			lowerIndex: 0,
			upperIndex: 7,
			output:     1,
		},
		{
			widthArray: []int{1, 2, 2, 2, 1},
			lowerIndex: 2,
			upperIndex: 3,
			output:     2,
		},
		{
			widthArray: []int{1, 2, 2, 2, 1},
			lowerIndex: 1,
			upperIndex: 4,
			output:     1,
		},
		{
			widthArray: []int{1, 2, 2, 2, 1},
			lowerIndex: 2,
			upperIndex: 4,
			output:     1,
		},
	}

	for _, test := range tests {
		if got := serviceLane(test.widthArray, test.lowerIndex, test.upperIndex); got != test.output {
			t.Errorf(
				"For widthArray=%v lowerIndex=%v upperIndex=%v; Got %v while Expecting %v",
				test.widthArray, test.lowerIndex, test.upperIndex, got, test.output,
			)
		}
	}
}
