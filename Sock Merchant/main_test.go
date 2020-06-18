package main

import "testing"


func TestCode(t *testing.T) {
	inputOne := []int {1, 2, 1, 2, 1, 3, 2}
	inputTwo := []int {10, 20, 20, 10, 10, 30, 50, 10, 20}
	var tests = []struct {
		input []int
		output int
	} {
		{inputOne, 2},
		{inputTwo, 3},
	}

	for _, test := range tests {
		if got := SockMerchant(test.input); got != test.output {
			t.Errorf("Staircase(%v) = %v", test.input, got)
		}
	}
}