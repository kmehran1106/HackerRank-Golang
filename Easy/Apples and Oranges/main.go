package main

import (
	"fmt"
)

func countAO(houseStart, houseEnd, apppleLoc, orangeLoc int, apples, oranges []int) (int, int) {
	a := 0
	o := 0

	for _, v := range apples {
		x := v + apppleLoc
		if houseStart <= x && x <= houseEnd {
			a++
		}
	}
	for _, v := range oranges {
		x := v + orangeLoc
		if houseStart <= x && x <= houseEnd {
			o++
		}
	}
	return a, o
}

func main() {
	fmt.Println(countAO(
		7, 11, 5, 15, []int{-2, 2, 1}, []int{5, -6}),
	)
}
