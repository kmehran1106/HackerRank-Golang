package main

import (
	"fmt"
	"sort"
)

func getSpaceStationMax(n int, arr []int) (output int) {
	sort.Ints(arr)

	f := arr[0]
	l := n - 1 - arr[len(arr)-1]
	if f > l {
		output = f
	} else {
		output = l
	}

	for i := 1; i < len(arr); i++ {
		d := (arr[i] - arr[i-1]) / 2
		if d > output {
			output = d
		}
	}
	return output
}

func main() {
	fmt.Println(getSpaceStationMax(5, []int{0, 4}))
}
