package main

import (
	"fmt"
)

func findDigits(x int) (r int) {
	y := x
	for y > 0 {
		n := y % 10
		if n != 0 && x%n == 0 {
			r++
		}
		y = y / 10
	}
	return r
}

func main() {
	fmt.Println(findDigits(1012))
}
