package main

import (
	"fmt"
)

func chocolateFeast(n, c, m int) (output int) {
	output = n / c
	wrappers := output

	for wrappers >= m {
		output += wrappers / m
		wrappers = wrappers/m + wrappers%m
	}
	return output
}

func main() {
	fmt.Println(chocolateFeast(7, 3, 2))
}
