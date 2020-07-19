package main

import "fmt"

func taumBirthday(b, w, bc, wc, z int) (output int) {
	var x, y int
	if bc > wc && z < bc-wc {
		x = w * wc
		y = b * (wc + z)
	} else if wc > bc && z < wc-bc {
		x = b * bc
		y = w * (bc + z)
	} else {
		x = b * bc
		y = w * wc
	}
	output = x + y
	return output
}

func main() {
	fmt.Println(taumBirthday(5, 9, 2, 3, 4))
}
