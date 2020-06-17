package main

import (
	"fmt"
	"sort"
)

func hurdleRace(jump int, hurdles []int) (doses int) {
	sort.Ints(hurdles)
	m := hurdles[len(hurdles)-1]
	if m < jump {
		doses = 0
	} else {
		doses = m - jump
	}
	return doses
}

func main() {
	fmt.Println(hurdleRace(4, []int{1, 6, 3, 5, 2}))
}
