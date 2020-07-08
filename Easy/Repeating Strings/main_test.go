package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		s      string
		n      int
		output int
	}{
		{
			s:      "a",
			n:      1000000000000,
			output: 1000000000000,
		},
		{
			s:      "aba",
			n:      10,
			output: 7,
		},
	}

	for _, test := range tests {
		if got := repeatingStrings(test.s, test.n); got != test.output {
			t.Errorf(
				"Repeating Substring for String=%v and N=%v; Got%v; Expected %v",
				test.s, test.n, got, test.output,
			)
		}
	}

}
