package main

import (
	"testing"
)


func TestCode(t *testing.T) {
	var tests = []struct {
		catA int
		catB int
		mouse int
		output string
	} {
		{
			catA: 1,
			catB: 2,
			mouse: 3,
			output: "Cat B",
		},
		{
			catA: 1,
			catB: 3,
			mouse: 2,
			output: "Mouse C",
		},
	}

	for _, test := range tests {
		if got := catsAndMouse(test.catA, test.catB, test.mouse); got != test.output {
			t.Errorf(
				"For Cat A=%v, Cat B=%v, Mouse=%v; Expected %v; Got %v", 
				test.catA, test.catB, test.mouse, test.output, got,
			)
		}
	}
}