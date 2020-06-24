package main

import (
	"fmt"
)

func gradingStudents(x int) (r int) {
	r = x
	if x >= 38 {
		t := x % 5
		if t > 2 {
			r = x + (5 - t)
		}
	}
	return r
}

func main() {
	fmt.Println(gradingStudents(74))
	fmt.Println(gradingStudents(76))
	fmt.Println(gradingStudents(33))
}
