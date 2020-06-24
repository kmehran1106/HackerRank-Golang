package main

import (
	"fmt"
)

func circularRotation(a []int, k int) []int {
	// x := []int{}
	k = k % len(a)
	a = append(a[len(a)-k:], a[:len(a)-k]...)
	return a
}

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(circularRotation(x, 2))
}
