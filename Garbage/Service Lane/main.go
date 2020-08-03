package main

import (
	"fmt"
)

func getMin(array []int) (output int) {
	if len(array) == 0 {
		output = 0
	} else {
		output = array[0]
		for _, v := range array {
			if v < output {
				output = v
			}
		}
	}
	return output
}

func serviceLane(widthArray []int, lowerIndex int, upperIndex int) (output int) {
	output = getMin(widthArray[lowerIndex : upperIndex+1])
	return output
}

func main() {
	fmt.Println(serviceLane([]int{2, 3, 1, 2, 3, 2, 3, 3}, 0, 3))
}
