package main

import (
	"fmt"
	"strings"
)

func bomberman(seconds int, grid []string) (output []string) {
	if seconds < 2 {
		output = grid
	} else if seconds%2 == 0 {
		var row string = strings.Repeat("O", len(grid[0]))
		for i := 0; i < len(grid); i++ {
			output = append(output, row)
		}
	} else {
		nGrid := [][]int{}
		for _, row := range grid {
			nRow := []int{}
			for _, c := range row {
				if string(c) == "." {
					nRow = append(nRow, 0)
				} else {
					nRow = append(nRow, 2)
				}
			}
			nGrid = append(nGrid, nRow)
		}

		R := len(grid)
		C := len(grid[0])

		t := 1

		for t < 4+seconds%4 {

			t++

			destroyed := [][]int{}
			for r := 0; r < R; r++ {
				for c := 0; c < C; c++ {
					if nGrid[r][c] > 0 {
						nGrid[r][c]--
					}
					if nGrid[r][c] == 0 {
						if t%2 == 0 {
							nGrid[r][c] = 3
						} else {
							destroyed = append(destroyed, []int{r, c})
							if r < R-1 {
								destroyed = append(destroyed, []int{r + 1, c})
							}
							if r > 0 {
								destroyed = append(destroyed, []int{r - 1, c})
							}
							if c < C-1 {
								destroyed = append(destroyed, []int{r, c + 1})
							}
							if c > 0 {
								destroyed = append(destroyed, []int{r, c - 1})
							}
						}
					}
				}
			}

			if len(destroyed) > 0 {
				for r, row := range nGrid {
					for c := range row {
						nGrid[r][c] = 2
					}
				}
				for _, v := range destroyed {
					nGrid[v[0]][v[1]] = 0
				}
			}
		}

		for _, row := range nGrid {
			var x string
			for _, c := range row {
				if c == 2 {
					x += "O"
				} else {
					x += "."
				}
			}
			output = append(output, x)
		}
	}
	return output
}

func main() {
	grid := []string{
		".......",
		"...O...",
		"....O..",
		".......",
		"OO.....",
		"OO.....",
	}

	output := bomberman(3, grid)
	for _, v := range output {
		fmt.Println(v)
	}
}
