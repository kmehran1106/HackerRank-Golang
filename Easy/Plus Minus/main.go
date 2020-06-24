package main

import (
	"fmt"
)

// PlusMinus :
func PlusMinus(array []int) []float64 {
	x := []float64{0, 0, 0}
	t := float64(len(array))
	for _, v := range array {
		if v > 0 {
			x[0] += 1 / t
		} else if v < 0 {
			x[1] += 1 / t
		} else {
			x[2] += 1 / t
		}
	}
	return x
}

func main() {
	fmt.Println(PlusMinus([]int{1, 0, -1}))
}
