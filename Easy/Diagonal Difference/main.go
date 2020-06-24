package main

import (
	"fmt"
	"math"
)

func diagonalDifference(matrix [][]int) (output int) {
	length := len(matrix) - 1
	var f, s int
	for i, a := range matrix {
		f += a[i]
		s += a[length-i]
	}
	output = int(math.Abs(float64(f - s)))
	return output
}

func main() {
	matrix := [][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	fmt.Println(diagonalDifference(matrix))
}
