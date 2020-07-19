package main

import "testing"

func TestCode(t *testing.T) {
	var tests = []struct {
		b      int
		w      int
		bc     int
		wc     int
		z      int
		output int
	}{
		{3, 5, 3, 4, 1, 29},
		{10, 10, 1, 1, 1, 20},
		{5, 9, 2, 3, 4, 37},
		{3, 6, 9, 1, 1, 12},
		{7, 7, 4, 2, 1, 35},
		{3, 3, 1, 9, 2, 12},
	}

	for _, test := range tests {
		if got := taumBirthday(test.b, test.w, test.bc, test.wc, test.z); got != test.output {
			t.Errorf(
				"For b=%v, w=%v, bc=%v, wc=%v, z=%v; Expected %v; Got %v",
				test.b, test.w, test.bc, test.wc, test.z, test.output, got,
			)
		}
	}
}
