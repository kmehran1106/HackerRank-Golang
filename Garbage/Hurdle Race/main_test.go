package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		hurdles []int
		jump    int
		output  int
	}{
		{hurdles: []int{1, 6, 3, 5, 2}, jump: 4, output: 2},
		{hurdles: []int{2, 3}, jump: 6, output: 0},
	}

	for _, test := range tests {
		if got := hurdleRace(test.jump, test.hurdles); got != test.output {
			t.Errorf(
				"For Hurdles %v with Current Jump Height %v; Expected %v; Got %v",
				test.jump, test.hurdles, test.output, got,
			)
		}
	}
}
