package main

import (
	"fmt"
)

func migratoryBirds(slice []int) (x int) {
	m := map[int]int{0: 0}
	x = 0
	for _, v := range slice {
		m[v]++
		if m[v] > m[x] || (m[v] == m[x] && v < x) {
			x=v
		}
	}
	return x
}

func main() {
	fmt.Println(migratoryBirds([]int{1, 1, 2, 2, 3}))
}

