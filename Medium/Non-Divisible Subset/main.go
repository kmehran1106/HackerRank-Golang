package main

import (
	"fmt"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getSubset(array []int, k int) (output int) {
	m := k / 2
	f := make([]int, k)

	for _, v := range array {
		d := v % k
		f[d]++
	}

	if k%2 == 0 {
		f[m] = min(f[m], 1)
	}
	output = min(f[0], 1)
	for i := 1; i <= m; i++ {
		output += max(f[i], f[k-i])
	}
	return output
}

func main() {
	array := []int{1, 7, 2, 4}
	k := 3
	fmt.Println(getSubset(array, k))
}
