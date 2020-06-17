package main

import (
	"fmt"
)

func divisibleSumPair(divisor int, numbers []int) (output int) {
	for i, v := range numbers {
		for _, x := range numbers[i+1:] {
			if (v+x)%divisor == 0 {
				output++
			}
		}
	}
	return output
}

func main() {
	fmt.Println(divisibleSumPair(3, []int{1, 3, 2, 6, 1, 2}))
}
