package main

import (
	"fmt"
)

func breakingRecords(scores []int) (min, max int) {
	top := scores[0]
	low := scores[0]

	for i := 1; i < len(scores); i++ {
		if scores[i] > top {
			top = scores[i]
			max++
		}
		if scores[i] < low {
			low = scores[i]
			min++
		}
	}
	return max, min
}

func main() {
	scores := []int{10, 5, 20, 20, 4, 5, 2, 25, 1}
	fmt.Println(breakingRecords(scores))
}
