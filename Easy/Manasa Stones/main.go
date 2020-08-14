package main

import (
	"fmt"
	"sort"
)

func manasaStones(n, a, b int) []int {
	m := map[int]bool{}
	for i := 0; i < n; i++ {
		x := a*(n-1-i) + b*i
		m[x] = true
	}
	s := []int{}
	for k := range m {
		s = append(s, k)
	}
	sort.Ints(s)
	return s
}

func main() {
	fmt.Println(manasaStones(4, 10, 100))
}
