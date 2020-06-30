package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		s      string
		t      string
		k      int
		output bool
	}{
		{s: "qwerasdf", t: "qwerbsdf", k: 6, output: false},
		{s: "hackerhappy", t: "hackerrank", k: 9, output: true},
		{s: "aba", t: "aba", k: 7, output: true},
		{s: "ashley", t: "ash", k: 2, output: false},
		{
			s:      "asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv",
			t:      "bsdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv",
			k:      100,
			output: false,
		},
	}

	for _, test := range tests {
		if got := AppendDelete(test.s, test.t, test.k); got != test.output {
			t.Errorf(
				"For s=%v t=%v k=%v; Expected Rounded Input %v; Got %v",
				test.s, test.t, test.k, test.output, got,
			)
		}
	}
}
