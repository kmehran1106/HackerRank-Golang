package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		n      int
		c      int
		m      int
		output int
	}{
		{7, 3, 2, 3},
		{10, 2, 5, 6},
		{12, 4, 4, 3},
		{6, 2, 2, 5},
	}

	for _, test := range tests {
		if got := chocolateFeast(test.n, test.c, test.m); got != test.output {
			t.Errorf(
				"For n=%v c=%v m=%v; Got %v while Expecting %v",
				test.n, test.c, test.m, got, test.output,
			)
		}
	}
}
