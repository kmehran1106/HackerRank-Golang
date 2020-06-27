package main

import (
	"fmt"
)

func sequenceEquation(px []int) (py []int) {
	d := map[int]int{}
	f := map[int]int{}
	for i, v := range px {
		d[i+1] = v
	}
	for k, v := range d {
		f[v] = k
	}
	for i := 0; i < len(px); i++ {
		py = append(py, f[f[i+1]])
	}
	return py
}

func main() {
	x := []int{2, 3, 1}
	fmt.Println(sequenceEquation(x))
}
