package main

import (
	"fmt"
	"strconv"
)

func getQueensAttack(n int, k int, rQ int, cQ int, obstacles [][]int) (output int) {
	/*
			Args:
		        n (int): [Chessboard Size Number. The Board will be nxn]
		        k (int): [Number of obstacles on Chessboard.]
		        rQ (int): [Row Number of Queen Position]
		        cQ (int): [Column Number of Queen Position]
		        obstacles ([][]int): [Array of Int Arrays denoting Obstacles Position]

		    Returns:
		        output (int): [Number of Places in the Board the Queen can Move to]
	*/
	obstacleMap := map[string]bool{}
	for _, v := range obstacles {
		obstacleMap[strconv.Itoa(v[0])+"_"+strconv.Itoa(v[1])] = true
	}
	fmt.Println(obstacleMap)

	// calculate top and bottom
	var i, j int

	i, j = rQ, cQ
	for i < n {
		if v := obstacleMap[strconv.Itoa(i+1)+"_"+strconv.Itoa(j)]; !v {
			output++
			i++
		} else {
			break
		}
	}
	fmt.Printf("Result After Top %v\n", output)

	i, j = rQ, cQ
	for i > 1 {
		if v := obstacleMap[strconv.Itoa(i-1)+"_"+strconv.Itoa(j)]; !v {
			output++
			i--
		} else {
			break
		}
	}
	fmt.Printf("Result After Bot %v\n", output)

	i, j = rQ, cQ
	for j < n {
		if v := obstacleMap[strconv.Itoa(i)+"_"+strconv.Itoa(j+1)]; !v {
			output++
			j++
		} else {
			break
		}
	}
	fmt.Printf("Result After Right %v\n", output)

	i, j = rQ, cQ
	for j > 1 {
		if v := obstacleMap[strconv.Itoa(i)+"_"+strconv.Itoa(j-1)]; !v {
			output++
			j--
		} else {
			break
		}
	}
	fmt.Printf("Result After Left %v\n", output)

	// Diagonals
	i, j = rQ, cQ
	for i < n && j < n {
		if v := obstacleMap[strconv.Itoa(i+1)+"_"+strconv.Itoa(j+1)]; !v {
			output++
			i++
			j++
		} else {
			break
		}
	}
	fmt.Printf("Result After Top-Right %v\n", output)

	i, j = rQ, cQ
	for i < n && j > 1 {
		if v := obstacleMap[strconv.Itoa(i+1)+"_"+strconv.Itoa(j-1)]; !v {
			output++
			i++
			j--
		} else {
			break
		}
	}
	fmt.Printf("Result After Top-Left %v\n", output)

	i, j = rQ, cQ
	for i > 1 && j > 1 {
		if v := obstacleMap[strconv.Itoa(i-1)+"_"+strconv.Itoa(j-1)]; !v {
			output++
			i--
			j--
		} else {
			break
		}
	}
	fmt.Printf("Result After Bot-Left %v\n", output)

	i, j = rQ, cQ
	for i > 1 && j < n {
		if v := obstacleMap[strconv.Itoa(i-1)+"_"+strconv.Itoa(j+1)]; !v {
			output++
			i--
			j++
		} else {
			break
		}
	}
	fmt.Printf("Result After Bot-Right %v\n", output)

	return output
}

func main() {
	n, k, rQ, cQ := 5, 3, 4, 3
	obstacles := [][]int{
		{5, 5},
		{4, 2},
		{2, 3},
	}
	fmt.Println(getQueensAttack(n, k, rQ, cQ, obstacles))
}
