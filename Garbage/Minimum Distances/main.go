package main

import (
	"fmt"
)

func minimumDistances(array []int) (output int) {
	r := 100000
	f := false
	m := map[int][]int{}

	for i, v := range array {
		_, ok := m[v]
		if !ok {
			m[v] = []int{i}
		} else {
			m[v] = append(m[v], i)
		}
	}

	for _, v := range m {
		if len(v) >= 2 {
			f = true
			t := v[1] - v[0]
			if t < r {
				r = t
			}
		}
	}
	if f {
		output = r
	} else {
		output = -1
	}
	return output
}

func main() {
	fmt.Println(minimumDistances([]int{7, 1, 3, 4, 1, 7}))
}
