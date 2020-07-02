package main

import (
	"fmt"
	"math"
)

// SherlockSquare :
func SherlockSquare(a, b int) (output int) {
	x := math.Floor(math.Sqrt(float64(b)))
	y := math.Ceil(math.Sqrt(float64(a)))
	return int(x) - int(y) + 1
}

func main() {
	fmt.Println(SherlockSquare(3, 9))
}
