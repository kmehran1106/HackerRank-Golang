package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		clouds []int
		k      int
		output int
	}{
		{
			clouds: []int{0, 0, 1, 0, 0, 1, 1, 0},
			k:      2,
			output: 92,
		},
		{
			clouds: []int{1, 1, 1, 0, 1, 1, 0, 0, 0, 0},
			k:      3,
			output: 80,
		},
	}

	for _, test := range tests {
		if got := jumpingCloud(test.clouds, test.k); got != test.output {
			t.Errorf(
				"For Clouds %v and Jump Length %v; Expected %v; Got %v",
				test.clouds, test.k, test.output, got,
			)
		}
	}
}
