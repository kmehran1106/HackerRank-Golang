package main

import (
	"fmt"
	"math"
	"strings"
)

func encryption(s string) (output string) {
	l := math.Sqrt(float64(len(s)))
	r := int(math.Floor(l))
	c := int(math.Ceil(l))

	for r*c < len(s) {
		if r < c {
			r++
		} else {
			c++
		}
	}
	array := []string{}

	for i := 0; i < r; i++ {
		first := i * c
		last := first + c
		if last > len(s) {
			last = len(s)
		}
		sub := s[first:last]
		for len(sub) < c {
			sub += " "
		}
		array = append(array, sub)
	}
	fmt.Println(array)

	for i := 0; i < c; i++ {
		t := ""
		for _, sub := range array {
			t += string(sub[i])
		}
		t = strings.TrimSpace(t)
		output += t + " "
	}
	output = strings.TrimSpace(output)
	return output
}

func main() {
	fmt.Println(encryption("feedthedog"))
}
