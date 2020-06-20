package main

import (
	"fmt"
)

func bonAppetit(array []int, index int, paid int) int {
	array = append(array[:index], array[index+1:]...)
	var actual float64
	for _, v := range array {
		actual += float64(v) / 2
	}
	return int(float64(paid) - actual)
}

func main() {
	fmt.Println(bonAppetit([]int{3, 10, 2, 9}, 1, 12))
	fmt.Println(bonAppetit([]int{3, 10, 2, 9}, 1, 7))
}
