package main

import "testing"

func TestCode(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"12:05:00AM", "00:05:00"},
		{"12:55:32PM", "12:55:32"},
		{"04:00:00AM", "04:00:00"},
		{"03:21:33PM", "15:21:33"},
	}

	for _, test := range tests {
		if got := TimeConversion(test.input); got != test.output {
			t.Errorf("TimeConversion(%v) = %v", test.input, got)
		}
	}
}
