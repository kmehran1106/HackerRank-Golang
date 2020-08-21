package main

import (
	"fmt"
	"math"
)

func surfaceArea(grid [][]int) int {
	area := 2 * len(grid) * len(grid[0])
	a := [][]int{}

	x := make([]int, len(grid[0])+2)
	a = append(a, x)
	for _, row := range grid {
		r := []int{}

		r = append(r, 0)
		r = append(r, row...)
		r = append(r, 0)

		a = append(a, r)
	}
	a = append(a, x)

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a[i]); j++ {
			area += int(math.Abs(float64(a[i][j] - a[i-1][j])))
			area += int(math.Abs(float64(a[i][j] - a[i][j-1])))
		}
	}
	return area
}

func main() {
	grid := [][]int{
		{1, 3, 4},
		{2, 2, 3},
		{1, 2, 4},
	}
	fmt.Println(surfaceArea(grid))
}
