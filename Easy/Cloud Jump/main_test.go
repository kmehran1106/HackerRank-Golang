package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		clouds []int
		output int
	}{
		{
			clouds: []int{0, 0, 1, 0, 0, 1, 0},
			output: 4,
		},
		{
			clouds: []int{0, 0, 0, 1, 0, 0},
			output: 3,
		},
	}

	for _, test := range tests {
		if got := cloudJump(test.clouds); got != test.output {
			t.Errorf(
				"For Clouds %v; Expected %v; Got %v",
				test.clouds, test.output, got,
			)
		}
	}
}
