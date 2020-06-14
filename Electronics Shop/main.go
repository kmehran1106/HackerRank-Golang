package main

import (
	"fmt"
	"sort"
)

func electronicsShop(budget int, keyboards []int, drives []int) (x int) {
	x = -1

	sort.Slice(keyboards, func(i, j int) bool { return keyboards[i] > keyboards[j] })
	sort.Slice(drives, func(i, j int) bool { return drives[i] < drives[j] })

	for _, k := range keyboards {
		for _, d := range drives {
			s := k + d
			if s > budget {
				break
			}
			if s > x {
				x = s
			}
		}
	}
	return x
}

func main() {
	fmt.Println(electronicsShop(5, []int{4}, []int{5}))
}