package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		d1     int
		m1     int
		y1     int
		d2     int
		m2     int
		y2     int
		output int
	}{
		{9, 6, 2015, 6, 6, 2015, 45},
		{2, 7, 1014, 1, 1, 1014, 3000},
		{2, 6, 2014, 7, 6, 2014, 0},
		{15, 7, 2014, 1, 7, 2015, 0},
	}

	for _, test := range tests {
		if got := libraryFine(test.d1, test.m1, test.y1, test.d2, test.m2, test.y2); got != test.output {
			t.Errorf(
				"For ReturnDate=%v-%v-%v and ExpectedDate=%v-%v-%v; Expected Rounded Input %v; Got %v",
				test.d1, test.m1, test.y1, test.d2, test.m2, test.y2, test.output, got,
			)
		}
	}
}
