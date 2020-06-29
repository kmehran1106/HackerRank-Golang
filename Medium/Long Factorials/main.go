package main

import (
	"fmt"
	"math/big"
)

func longFactorials(n int) *big.Int {
	var x = new(big.Int)
	x.MulRange(1, int64(n))
	return x
}

func main() {
	fmt.Println(longFactorials(25))
}
