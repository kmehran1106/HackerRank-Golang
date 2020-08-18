package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		b      string
		output bool
	}{
		{"RBY_YBR", true},
		{"X_Y__X", false},
		{"__", true},
		{"B_RRBR", true},
		{"AABBC", false},
		{"AABBC_C", true},
		{"_", true},
		{"DD__FQ_QQF", true},
		{"AABCBC", false},
	}

	for _, test := range tests {
		if got := happyLadybugs(test.b); got != test.output {
			t.Errorf("happyLadybugs(%v); Got%v;Expecting %v", test.b, got, test.output)
		}
	}
}
