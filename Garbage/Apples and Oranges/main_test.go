package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		houseStart int
		houseEnd   int
		appleLoc   int
		orangeLoc  int
		apples     []int
		oranges    []int
		appleOut   int
		orangeOut  int
	}{
		{
			houseStart: 7,
			houseEnd:   11,
			appleLoc:   5,
			orangeLoc:  15,
			apples:     []int{-2, 2, 1},
			oranges:    []int{5, -6},
			appleOut:   1,
			orangeOut:  1,
		},
	}

	for _, test := range tests {
		a, o := countAO(
			test.houseStart, test.houseEnd, test.appleLoc, test.orangeLoc, test.apples, test.oranges,
		)
		if a != test.appleOut {
			t.Errorf(" Got Apples %v while Expecting %v", a, test.appleOut)
		}
		if o != test.orangeOut {
			t.Errorf("Got Oranges %v while Expecting %v", o, test.orangeOut)
		}
	}
}
