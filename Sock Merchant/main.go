package main

import "fmt"


func SockMerchant(array []int) (pairs int) {
	sockMap := make(map[int]int)

	for _, s := range array {
		v, ok := sockMap[s]
		if !ok {
			sockMap[s] = 1
		} else {
			sockMap[s] = v + 1
		}
	}

	for _, val := range sockMap {
		pairs += val / 2
	}
	return pairs
}

func main() {
	s := []int{1, 2, 1, 2}
	fmt.Println(SockMerchant(s))
}