package main

import "fmt"

func gcdPair(a, b int) int {
	for b > 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func gcdList(x []int) int {
	r := x[0]
	for _, v := range x[1:] {
		r = gcdPair(r, v)
	}
	return r
}

func lcmPair(a, b int) int {
	return a * (b / gcdPair(a, b))
}

func lcmList(x []int) int {
	r := x[0]
	for _, v := range x[1:] {
		r = lcmPair(r, v)
	}
	return r
}

func betweenTwoSets(first, second []int) int {
	count := 0
	lcm := lcmList(first)
	gcd := gcdList(second)
	i, j := lcm, 2

	for i <= gcd {
		if gcd%i == 0 {
			count++
		}
		i = lcm * j
		j++
	}
	return count
}

func main() {
	fmt.Println(betweenTwoSets([]int{2, 4}, []int{16, 32, 96}))
}
