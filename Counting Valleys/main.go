package main

import (
	"fmt"
)

func countingValleys(steps []string) (output int) {
	height := 0
	count := 0
	inValley := false

	for _, v := range steps {
		if string(v) == "D" {
			height--
		} else if string(v) == "U" {
			height++
		}

		if height < 0 && !inValley {
			count++
			inValley = true
		}
		if height >= 0 && inValley {
			inValley = false
		}
	}
	return count
}

func main() {
	input := []string{"U", "D", "D", "D", "U", "D", "U", "U"}
	fmt.Println(countingValleys(input))
}
