package main

import (
	"fmt"
)

func isCancelled(threshold int, timings []int) bool {
	count := 0
	for _, v := range timings {
		if v <= 0 {
			count++
		}
	}
	return count < threshold
}

func main() {
	scores := []int{-1, -3, 4, 2}
	fmt.Println(isCancelled(3, scores))
}
