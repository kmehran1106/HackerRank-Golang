package staircase

import "testing"


func TestCode(t *testing.T) {
	var tests = []struct {
		input int
		output string
	} {
		{2, " #\n##"},
		{3, "  #\n ##\n###"},
	}

	for _, test := range tests {
		if got := Staircase(test.input); got != test.output {
			t.Errorf("Staircase(%v) = %v", test.input, got)
		}
	}
}