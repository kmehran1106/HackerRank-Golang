package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		n      int
		k      int
		output []int
	}{
		{2, 1, []int{2, 1}},
		{3, 0, []int{1, 2, 3}},
		{3, 2, []int{-1}},
		{10, 1, []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9}},
	}

	for _, test := range tests {
		got := absolutePermutation(test.n, test.k)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf("Absolute Permutation Function Got%v while Expecting %v", got, test.output)
			}
		}
	}
}
