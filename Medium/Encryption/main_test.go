package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{
			input:  "feedthedog",
			output: "fto ehg ee dd",
		},
		{
			input:  "haveaniceday",
			output: "hae and via ecy",
		},
		{
			input:  "chillout",
			output: "clu hlt io",
		},
	}

	for _, test := range tests {
		if got := encryption(test.input); got != test.output {
			t.Errorf(
				"Organzing Containers for Input=%v; Got%v; Expected %v",
				test.input, got, test.output,
			)
		}
	}

}
