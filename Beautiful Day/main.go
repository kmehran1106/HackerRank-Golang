package main

import (
	"fmt"
	"math"
)

func reverse(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}

func beautifulDay(start, end, divisor int) int {
	count := 0

	for x := start; x <= end; x++ {
		s := math.Abs(float64(x - reverse(x)))
		if int(s)%divisor == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(beautifulDay(20, 23, 6))
}
