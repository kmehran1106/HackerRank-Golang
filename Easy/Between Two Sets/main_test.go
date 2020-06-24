package main

import (
	"testing"
)

func TestGCDPair(t *testing.T) {
	var tests = []struct {
		a      int
		b      int
		output int
	}{
		{8, 40, 8},
		{11, 17, 1},
	}

	for _, test := range tests {
		if got := gcdPair(test.a, test.b); got != test.output {
			t.Errorf("GCD of %v and %v; Expected %v; Got %v", test.a, test.b, test.output, got)
		}
	}
}

func TestGCDList(t *testing.T) {
	var tests = []struct {
		x      []int
		output int
	}{
		{[]int{8, 40, 160}, 8},
	}

	for _, test := range tests {
		if got := gcdList(test.x); got != test.output {
			t.Errorf("GCD of %v; Expected %v; Got %v", test.x, test.output, got)
		}
	}
}

func TestLCMPair(t *testing.T) {
	var tests = []struct {
		a      int
		b      int
		output int
	}{
		{8, 40, 40},
		{11, 17, 187},
	}

	for _, test := range tests {
		if got := lcmPair(test.a, test.b); got != test.output {
			t.Errorf("LCM of %v and %v; Expected %v; Got %v", test.a, test.b, test.output, got)
		}
	}
}

func TestLCMList(t *testing.T) {
	var tests = []struct {
		x      []int
		output int
	}{
		{[]int{8, 40, 160}, 160},
	}

	for _, test := range tests {
		if got := lcmList(test.x); got != test.output {
			t.Errorf("LCM of %v; Expected %v; Got %v", test.x, test.output, got)
		}
	}
}

func TestCode(t *testing.T) {
	var tests = []struct {
		first  []int
		second []int
		output int
	}{
		{first: []int{2, 4}, second: []int{16, 32, 96}, output: 3},
	}

	for _, test := range tests {
		if got := betweenTwoSets(test.first, test.second); got != test.output {
			t.Errorf("For %v and %v; Expected %v; Got %v", test.first, test.second, test.output, got)
		}
	}
}
