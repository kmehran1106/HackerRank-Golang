package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		divisor int
		numbers []int
		output  int
	}{
		{divisor: 3, numbers: []int{1, 3, 2, 6, 1, 2}, output: 5},
	}

	for _, test := range tests {
		if got := divisibleSumPair(test.divisor, test.numbers); got != test.output {
			t.Errorf(
				"For Numbers %v and Divisor %v; Expected Output %v; Got %v",
				test.numbers, test.divisor, test.output, got,
			)
		}
	}
}
