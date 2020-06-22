package main

import (
	"fmt"
)

func birthdayBar(array []int, date int, month int) int {
	temp := array[:len(array)-month+1]
	result := 0

	for i, v := range temp {
		x := 1
		s := v
		for x < month {
			s += array[i+x]
			x++
		}
		if s == date {
			result++
		}
	}
	return result
}

func main() {
	x := []int{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}
	y := []int{1, 2, 1, 3, 2}
	fmt.Println(birthdayBar(x, 18, 7))
	fmt.Println(birthdayBar(y, 3, 2))
}
