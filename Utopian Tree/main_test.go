package main

import "testing"

func TestCode(t *testing.T) {
	var tests = []struct {
		input  int
		output int
	}{
		{4, 7},
		{5, 14},
	}

	for _, test := range tests {
		if got := utopianTree(test.input); got != test.output {
			t.Errorf("For Height %v; Expected %v; Got %v", test.input, test.output, got)
		}
	}
}
