package main

import (
	"fmt"
	"sort"
)

func getRowSum(array [][]int) []int {
	output := make([]int, len(array))
	for index, sub := range array {
		for _, v := range sub {
			output[index] += v
		}
	}
	return output
}

func getColSum(array [][]int) []int {
	output := make([]int, len(array[0]))
	for i, sub := range array {
		for j := range sub {
			output[i] += array[j][i]
		}
	}
	return output
}

func organizingContainers(array [][]int) (output bool) {
	output = true
	rowTotals := getRowSum(array)
	sort.Sort(sort.IntSlice(rowTotals))
	colTotals := getColSum(array)
	sort.Sort(sort.IntSlice(colTotals))

	for index, value := range rowTotals {
		if value != colTotals[index] {
			output = false
			break
		}
	}
	return output
}

func main() {
	array := [][]int{
		{1, 3, 1},
		{2, 1, 2},
		{3, 3, 3},
	}
	fmt.Println(organizingContainers(array))
}
