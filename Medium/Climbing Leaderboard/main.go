package main

import (
	"fmt"
	"sort"
)

func unique(slice []int) (list []int) {
	keys := make(map[int]bool)
	for _, entry := range slice {
		if _, ok := keys[entry]; !ok {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func binarySearch(array []int, left int, right int, x int) (index int) {
	if right >= left {
		mid := (right + left) / 2
		if array[mid] == x {
			index = mid
		} else if array[mid] > x {
			return binarySearch(array, left, mid-1, x)
		} else {
			return binarySearch(array, mid+1, right, x)
		}
	} else {
		index = right
	}
	return index
}

func climbingLeaderboardBinary(scores, alice []int) (ranks []int) {
	scores = unique(scores)
	sort.Ints(scores)
	n := len(scores)

	for _, v := range alice {
		index := binarySearch(scores, 0, n-1, v)
		ranks = append(ranks, n-index)
	}
	return ranks
}

func climbingLeaderboardIterative(scores, alice []int) (ranks []int) {
	scores = unique(scores)
	sort.Ints(scores)
	n := len(scores)

	index := 0
	for _, v := range alice {
		for n > index && v >= scores[index] {
			index++
		}
		ranks = append(ranks, n+1-index)
	}
	return ranks
}

func main() {
	scores := []int{100, 90, 90, 80, 75, 60}
	alice := []int{50, 65, 77, 90, 102}
	fmt.Println(climbingLeaderboardIterative(scores, alice))
	fmt.Println(climbingLeaderboardBinary(scores, alice))
}
