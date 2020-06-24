package main

import (
	"fmt"
)

func sum(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}

func viralAdvertising(days int) int {
	l := make([]int, days)
	l[0] = 2
	for i := 1; i < days; i++ {
		l[i] = (l[i-1] * 3) / 2
	}
	return sum(l)
}

func main() {
	fmt.Println(viralAdvertising(5))
	fmt.Println(viralAdvertising(3))
}
