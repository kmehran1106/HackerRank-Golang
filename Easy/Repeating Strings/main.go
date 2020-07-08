package main

import (
	"fmt"
)

func repeatingStrings(s string, n int) (output int) {
	full := n / len(s)
	extras := n % len(s)
	var c1, cf, ce int
	for _, v := range s {
		if string(v) == "a" {
			c1++
		}
	}
	cf = c1 * full
	if extras != 0 {
		for _, v := range s[:extras] {
			if string(v) == "a" {
				ce++
			}
		}
	}
	output = cf + ce
	return output
}

func main() {
	fmt.Println(repeatingStrings("aba", 10))
}
