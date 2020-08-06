package main

import (
	"fmt"
)

func getWorkbookSpecialProblems(n int, k int, arr []int) (output int) {
	page := 1
	for ch := 0; ch < n; ch++ {
		for p := 1; p <= arr[ch]; p++ {
			if p == page {
				output++
			}
			if p%k == 0 || p == arr[ch] {
				page++
			}
		}
	}
	return output
}

func main() {
	fmt.Println(getWorkbookSpecialProblems(5, 3, []int{4, 2, 6, 1, 10}))
}
