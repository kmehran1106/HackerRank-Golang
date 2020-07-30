package main

import (
	"fmt"
)

func halloweenSale(p, d, m, s int) (output int) {
	if s < p {
		return 0
	}
	s = s - p
	output++
	for s > 0 {
		if p-d > m {
			p = p - d
		} else {
			p = m
		}
		if s-p < 0 {
			break
		}
		s = s - p
		output++
	}
	return output
}

func main() {
	fmt.Println(halloweenSale(20, 3, 6, 80))
	fmt.Println(halloweenSale(20, 3, 6, 85))
}
