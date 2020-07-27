package main

import (
	"fmt"
)

func cloudJump(clouds []int) (output int) {
	var i int
	last := len(clouds) - 1
	for i < last-2 {
		if clouds[i+2] == 0 {
			i += 2
		} else {
			i++
		}
		output++
	}
	output++
	return output
}

func main() {
	c := []int{0, 0, 1, 0, 0, 1, 0}
	fmt.Println(cloudJump(c))
}
