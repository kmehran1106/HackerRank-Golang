package main

import (
	"fmt"
	"math"
)

func drawingBook(total, page int) (turns int) {
	if page == 1 || (total%2 == 0 && page == total) || (total%2 != 0 && page >= total-1) {
		turns = 0
	} else {
		var x float64
		if page <= total/2 {
			x = float64(page-1) / 2.0
			x = math.Ceil(x)
		} else if total%2 != 0 {
			x = float64(total-page) / 2.0
		} else {
			x = float64(total-page) / 2.0
			x = math.Ceil(x)
		}
		turns = int(x)
	}

	return turns
}

func main() {
	fmt.Println(drawingBook(6, 2))
	fmt.Println(drawingBook(5, 4))
	fmt.Println(drawingBook(70809, 46090))
	fmt.Println(drawingBook(6, 5))
}
