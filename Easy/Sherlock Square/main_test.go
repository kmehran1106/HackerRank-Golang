package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		a      int
		b      int
		output int
	}{
		{3, 9, 2},
		{17, 24, 0},
	}

	for _, test := range tests {
		if got := SherlockSquare(test.a, test.b); got != test.output {
			t.Errorf(
				"For Lower=%v and Upper=%v; Expected Rounded Input %v; Got %v",
				test.a, test.b, test.output, got,
			)
		}
	}
}
