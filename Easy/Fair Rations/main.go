package main

import (
	"fmt"
)

func fairRations(array []int) (output int) {
	sum := 0
	for _, v := range array {
		sum += v
	}
	if sum%2 != 0 {
		return 0
	}

	for i := 0; i < len(array)-1; i++ {
		if array[i]%2 != 0 {
			array[i] = array[i] + 1
			array[i+1] = array[i+1] + 1
			output += 2
		}
	}
	return output
}

func main() {
	fmt.Println(fairRations([]int{2, 3, 4, 5, 6}))
}
