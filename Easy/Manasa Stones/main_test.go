package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		n      int
		a      int
		b      int
		output []int
	}{
		{
			n: 3, a: 1, b: 2,
			output: []int{2, 3, 4},
		},
		{
			n: 4, a: 10, b: 100,
			output: []int{30, 120, 210, 300},
		},
	}

	for _, test := range tests {
		got := manasaStones(test.n, test.a, test.b)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf(
					"For Stones=%v DiffA=%v DiffB=%v; Got %v while expecting %v",
					test.n, test.a, test.b, got, test.output,
				)
			}
		}
	}
}
