package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		p      int
		q      int
		output string
	}{
		{1, 100, "1 9 45 55 99"},
		{100, 300, "297"},
		{400, 700, ""},
	}
	for _, test := range tests {
		if got := modifiedKaprekarNumber(test.p, test.q); got != test.output {
			t.Errorf(
				"For Range %v to %v for Modified Kaprekar Function; Got%v while Expecting %v",
				test.p, test.q, got, test.output,
			)
		}
	}

}
