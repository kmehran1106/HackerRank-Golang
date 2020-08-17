package main

import (
	"fmt"
)

func gridSearch(grid, pattern []string) bool {
	gRow := len(grid) - 1
	gCol := len(grid[0]) - 1

	pRow := len(pattern) - 1
	pCol := len(pattern[0]) - 1

	c := 0

	for i := 0; i <= gRow-pRow; i++ {
		for j := 0; j <= gCol-pCol; j++ {
			if grid[i][j:j+pCol+1] == pattern[0] {
				for k := 1; k < len(pattern); k++ {
					if grid[i+k][j:j+pCol+1] == pattern[k] {
						c++
						if c == len(pattern)-1 {
							return true
						}
					} else {
						c = 0
					}
				}
			}
		}
	}
	return false
}

func main() {
	g := []string{
		"7283455864",
		"6731158619",
		"8988242643",
		"3830589324",
		"2229505813",
		"5633845374",
		"6473530293",
		"7053106601",
		"0834282956",
		"4607924137",
	}
	p := []string{
		"9505",
		"3845",
		"3530",
	}
	fmt.Println(gridSearch(g, p))
}
