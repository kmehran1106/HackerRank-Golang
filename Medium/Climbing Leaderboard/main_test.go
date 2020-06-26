package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		scores []int
		alice  []int
		output []int
	}{
		{
			scores: []int{100, 90, 90, 80, 75, 60},
			alice:  []int{50, 65, 77, 90, 102},
			output: []int{6, 5, 4, 2, 1},
		},
		{
			scores: []int{100, 100, 50, 40, 40, 20, 10},
			alice:  []int{5, 25, 50, 120},
			output: []int{6, 4, 2, 1},
		},
	}

	for _, test := range tests {
		got := climbingLeaderboardIterative(test.scores, test.alice)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf("climbingLeaderboard Got%v while Expecting %v", got, test.output)
			}
		}
	}

	for _, test := range tests {
		got := climbingLeaderboardBinary(test.scores, test.alice)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf("climbingLeaderboard Got%v while Expecting %v", got, test.output)
			}
		}
	}

}
