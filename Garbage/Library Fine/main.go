package main

import (
	"fmt"
)

const fineY = 10000
const fineM = 500
const fineD = 15

func libraryFine(d1, m1, y1, d2, m2, y2 int) (output int) {
	if y1 > y2 {
		output = fineY
	} else if y1 < y2 {
		output = 0
	} else {
		if m1 > m2 {
			output = fineM * (m1 - m2)
		} else if m1 < m2 {
			output = 0
		} else {
			if d1 > d2 {
				output = fineD * (d1 - d2)
			} else {
				output = 0
			}
		}
	}
	return output
}

func main() {
	fmt.Println(libraryFine(9, 6, 2015, 6, 6, 2015))
}
