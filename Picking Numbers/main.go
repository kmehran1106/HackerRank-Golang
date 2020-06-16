package main

import (
	"fmt"
	"sort"
)

func pickingNumbers(array []int) (m int) {
	hashMap := make(map[int]int)
	sort.Ints(array)

	for _, v := range array {
		hashMap[v]++
	}
	for k, v := range hashMap {
		if v == -1 {
			continue
		}
		x := hashMap[k+1]
		y := x + v
		if y > m {
			m = y
			hashMap[k] = -1
		}
		if x > 0 {
			hashMap[k+1] = -1
		}
	}
	return m
}

func main() {
	fmt.Println(pickingNumbers([]int{1, 1, 2, 2, 1}))
}
