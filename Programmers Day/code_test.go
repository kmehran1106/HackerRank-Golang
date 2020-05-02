package programmers_day

import "testing"


func TestCode(t *testing.T) {
	var tests = []struct {
		input int
		output string
	} {
		{2017, "13.09.2017"},
		{2016, "12.09.2016"},
		{1800, "12.09.1800"},
		{1918, "26.09.1918"},
	}

	for _, test := range tests {
		if got := ProgrammersDay(test.input); got != test.output {
			t.Errorf("Staircase(%v) = %v", test.input, got)
		}
	}
}