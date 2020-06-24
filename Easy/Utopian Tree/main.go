package main

import "fmt"

func utopianTree(n int) (h int) {
	for i := 0; i <= n; i++ {
		if i%2 == 1 {
			h *= 2
		} else {
			h++
		}
	}
	return h
}

func main() {
	fmt.Println(utopianTree(4))
}
