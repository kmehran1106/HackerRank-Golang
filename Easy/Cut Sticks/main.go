package main

import (
	"fmt"
	"sort"
)

func cutSticks(input []int) (output []int) {
	l := len(input)
	p := -1
	sort.Ints(input)

	for _, v := range input {
		if v > p {
			output = append(output, l)
		}
		l--
		p = v
	}
	return output
}

func main() {
	input := []int{5, 4, 4, 2, 2, 8}
	fmt.Println(cutSticks(input))
}
