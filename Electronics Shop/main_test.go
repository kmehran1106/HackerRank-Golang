package main

import (
	"testing"
)


func TestCode(t *testing.T) {
	var tests = []struct {
		budget int
		keyboards []int
		drives []int
		output int
	} {
		{
			budget: 10,
			keyboards: []int{3, 1},
			drives: []int{5, 2, 8},
			output: 9,
		},
	}

	for _, test := range tests {
		if got := electronicsShop(test.budget, test.keyboards, test.drives); got != test.output {
			t.Errorf(
				"For Budget=%v, Keyboards=%v, Drives=%v; Expected %v; Got %v", 
				test.budget, test.keyboards, test.drives, test.output, got,
			)
		}
	}
}