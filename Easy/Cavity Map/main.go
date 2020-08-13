package main

import (
	"fmt"
)

func cavityMap(grid []string) []string {
	for i := 1; i < len(grid)-1; i++ {
		n := []byte{}
		n = append(n, grid[i][0])
		for j := 1; j < len(grid)-1; j++ {
			top := grid[i-1][j]
			bottom := grid[i+1][j]
			left := grid[i][j-1]
			right := grid[i][j+1]
			x := grid[i][j]
			if x > top && x > bottom && x > left && x > right {
				n = append(n, []byte("X")[0])
			} else {
				n = append(n, x)
			}
		}
		n = append(n, grid[i][len(grid)-1])
		grid[i] = string(n)
	}
	return grid
}

func main() {
	grid := []string{"1112", "1912", "1892", "1234"}
	fmt.Println(cavityMap(grid))
}
