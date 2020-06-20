package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		start   int
		end     int
		divisor int
		output  int
	}{
		{20, 23, 6, 2},
	}

	for _, test := range tests {
		if got := beautifulDay(test.start, test.end, test.divisor); got != test.output {
			t.Errorf(
				"Given Start %v and End %v For Divisor %v \n; Expected %v; Got %v",
				test.start, test.end, test.divisor, test.output, got,
			)
		}
	}
}
