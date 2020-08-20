package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		t      int
		output int
	}{
		{17, 5},
		{4, 6},
	}

	for _, test := range tests {
		if got := strangeCounter(test.t); got != test.output {
			t.Errorf("strangeCounter(%v); Got%v;Expecting %v", test.t, got, test.output)
		}
	}
}
