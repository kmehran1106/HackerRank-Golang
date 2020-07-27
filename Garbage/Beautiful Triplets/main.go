package main

import (
	"fmt"
)

func beautifulTriplets(d int, array []int) (output int) {
	m := map[int]bool{}
	for _, v := range array {
		m[v] = true
	}
	for _, i := range array {
		j := m[i+d]
		k := m[i+2*d]
		if j && k {
			output++
		}
	}
	return output
}

func main() {
	fmt.Println(beautifulTriplets(3, []int{1, 2, 4, 5, 7, 8, 10}))
}
