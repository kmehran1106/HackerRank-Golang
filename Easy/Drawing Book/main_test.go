package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	var tests = []struct {
		total int
		page  int
		turns int
	}{
		{total: 6, page: 2, turns: 1},
		{total: 5, page: 4, turns: 0},
		{total: 70809, page: 46090, turns: 12359},
		{total: 6, page: 5, turns: 1},
	}

	for _, test := range tests {
		if got := drawingBook(test.total, test.page); got != test.turns {
			t.Errorf(
				"For Total Pages %v and Input Page %v; Expected Turns %v; Got %v",
				test.total, test.page, test.turns, got,
			)
		}
	}
}
