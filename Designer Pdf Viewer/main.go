package main

import (
	"fmt"
)

func pdfViewer(heights []int, input string) (output int) {
	diff := 97
	max := 0
	for _, v := range input {
		m := heights[int(v) - diff]
		if m > max {
			max = m
		}
	}
	output = max * len(input)
	return output
}

func main() {
	heights := []int{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	input := "abc"
	fmt.Println(pdfViewer(heights, input))
}
