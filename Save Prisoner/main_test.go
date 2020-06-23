package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		prisoners int
		sweets    int
		chair     int
		output    int
	}{
		{
			prisoners: 7,
			sweets:    19,
			chair:     2,
			output:    6,
		},
		{
			prisoners: 13,
			sweets:    140874526,
			chair:     1,
			output:    13,
		},
	}

	for _, test := range tests {
		if got := savePrisoner(test.prisoners, test.sweets, test.chair); got != test.output {
			t.Errorf(
				"For Prisoners %v, Sweets %v, Start Chair %v; Expected %v; Got %v",
				test.prisoners, test.sweets, test.chair, test.output, got,
			)
		}
	}
}
