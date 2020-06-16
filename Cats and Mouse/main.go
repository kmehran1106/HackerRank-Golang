package main

import (
	"fmt"
	"math"
)

func catsAndMouse(x, y, z int) (v string) {
	catA := math.Abs(float64(x - z))
	catB := math.Abs(float64(y - z))
	if catA < catB {
		return "Cat A"
	} else if catB < catA {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main() {
	fmt.Println(catsAndMouse(1, 2, 3))
}
