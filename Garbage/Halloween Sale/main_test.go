package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		p      int
		m      int
		d      int
		s      int
		output int
	}{
		{20, 3, 6, 80, 6},
		{20, 3, 6, 85, 7},
	}

	for _, test := range tests {
		if got := halloweenSale(test.p, test.m, test.d, test.s); got != test.output {
			t.Errorf(
				"For p=%v m=%v d=%v s=%v; Got %v while Expecting %v",
				test.p, test.m, test.d, test.s, got, test.output,
			)
		}
	}
}
