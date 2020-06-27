package main

import (
	"fmt"
)

func jumpingCloud(clouds []int, k int) (e int) {
	e = 100
	n := len(clouds)
	// initial declarations
	i := k % n
	e -= (1 + clouds[i]*2)

	for i != 0 {
		i = (i + k) % n
		e -= (1 + clouds[i]*2)
	}
	return e
}

func main() {
	c := []int{0, 0, 1, 0, 0, 1, 1, 0}
	fmt.Println(jumpingCloud(c, 2))
}
