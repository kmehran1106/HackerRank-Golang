package main

import (
	"math/big"
	"testing"
)

func TestCode(t *testing.T) {
	x, _ := new(big.Int).SetString("15511210043330985984000000", 0)
	y, _ := new(big.Int).SetString("119622220865480194561963161495657715064383733760000000000", 0)
	var tests = []struct {
		input  int
		output *big.Int
	}{
		{input: 25, output: x},
		{input: 45, output: y},
	}

	for _, test := range tests {
		if got := longFactorials(test.input); got.Cmp(test.output) != 0 {
			t.Errorf(
				"For Input %v; Expected Rounded Input %v; Got %v",
				test.input, test.output, got,
			)
		}
	}
}
