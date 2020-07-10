package main

import (
	"fmt"
)

func equalizeArray(array []int) (output int) {
	numMap := map[int]int{}
	maxFreq := 0
	for _, v := range array {
		numMap[v]++
		if c := numMap[v]; c > maxFreq {
			maxFreq = c
		}
	}
	output = len(array) - maxFreq
	return output
}

func main() {
	fmt.Println(equalizeArray([]int{1, 2, 3, 1, 2, 3, 3, 3}))
}
