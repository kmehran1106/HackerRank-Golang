package main

import (
	"fmt"
)

func absolutePermutation(n, k int) (output []int) {
	a := make(map[int]bool)

	for i := 1; i <= n; i++ {
		a[i] = true
	}

	if k == 0 {
		for i := 1; i <= n; i++ {
			output = append(output, i)
		}
	} else if (n/k)%2 != 0 {
		output = append(output, -1)
	} else {
		condition := true
		for i := 1; i <= n; i++ {
			if condition {
				output = append(output, i+k)
			} else {
				output = append(output, i-k)
			}
			if i%k == 0 {
				condition = !condition
			}
		}
	}
	return output
}

func main() {
	fmt.Println(absolutePermutation(10, 1))
}
