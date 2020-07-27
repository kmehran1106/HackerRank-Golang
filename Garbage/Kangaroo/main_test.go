package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		x1     int
		v1     int
		x2     int
		v2     int
		output bool
	}{
		{x1: 0, v1: 3, x2: 4, v2: 2, output: true},
		{x1: 28, v1: 8, x2: 96, v2: 2, output: false},
	}

	for _, test := range tests {
		if got := kangaroo(test.x1, test.v1, test.x2, test.v2); got != test.output {
			t.Errorf(
				"x1=%v v1=%v x2=%v v2=%v; Expected %v; Got %v",
				test.x1, test.v1, test.x2, test.v2, test.output, got,
			)
		}
	}
}
