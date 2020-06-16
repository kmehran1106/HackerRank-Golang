package main

import (
	"fmt"
)

func kangaroo(x1, v1, x2, v2 int) bool {
	if x1 < x2 && v1 <= v2 {
		return false
	}
	n1 := float64(x2-x1) / float64(v1-v2)
	n2 := float64((x2 - x1) / (v1 - v2))

	if n1 == n2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2))
}
