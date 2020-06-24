package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		array  []int
		date   int
		month  int
		output int
	}{
		{
			array:  []int{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1},
			date:   18,
			month:  7,
			output: 3,
		},
		{
			array:  []int{1, 2, 1, 3, 2},
			date:   3,
			month:  2,
			output: 2,
		},
	}

	for _, test := range tests {
		if got := birthdayBar(test.array, test.date, test.month); got != test.output {
			t.Errorf(
				"Given Array %v, Date %v, Month %v\n; Expected %v; Got %v",
				test.array, test.date, test.month, test.output, got,
			)
		}
	}
}
