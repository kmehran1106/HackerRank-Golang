package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		array  []int
		k      int
		output int
	}{
		{
			array:  []int{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436},
			k:      7,
			output: 11,
		},
		{
			array:  []int{1, 7, 2, 4},
			k:      3,
			output: 3,
		},
	}

	for _, test := range tests {
		if got := getSubset(test.array, test.k); got != test.output {
			t.Errorf(
				"Non Divisible Subset for Array=%v and K=%v; Got%v; Expected %v",
				test.array, test.k, got, test.output,
			)
		}
	}

}
