package main

import (
	"fmt"
	"sort"
)

// MinMaxSum :
func MinMaxSum(array []int64) []int64{
	x := []int64{0, 0}
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })

	low := array[:len(array)-1]
	high := array[1:]
	
	for _, v := range(low) {
		x[0] += v
	}
	for _, v := range(high) {
		x[1] += v
	}
	return x
}

func main() {
	fmt.Println(MinMaxSum([]int64{1,2,5,4,3}))
}