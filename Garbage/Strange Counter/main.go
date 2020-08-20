package main

import (
	"fmt"
)

const startTime = 1
const startValue = 3

func strangeCounter(t int) int {
	x, y := startTime, startValue
	for t > y {
		x += y
		y *= 2
		if t < 2*y {
			break
		}
	}
	return y - (t - x)
}

func main() {
	fmt.Println(strangeCounter(17))
}
