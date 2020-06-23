package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		a      []int
		k      int
		output []int
	}{
		{
			a:      []int{1, 2, 3, 4},
			k:      2,
			output: []int{3, 4, 1, 2},
		},
	}

	for _, test := range tests {
		got := circularRotation(test.a, test.k)
		for i, v := range got {
			if v != test.output[i] {
				t.Errorf("circularRotation Got%v while Expecting %v", got, test.output)
			}
		}
	}

}
