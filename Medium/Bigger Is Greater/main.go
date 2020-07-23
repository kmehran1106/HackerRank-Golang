package main

import (
	"fmt"
)

func biggerIsGreater(s string) (output string) {
	var pivot int
	var i, j int = len(s) - 1, len(s) - 1
	array := []rune(s)

	// calculate the suffix
	for ; i > 0 && array[i-1] >= array[i]; i-- {
	}
	if i <= 0 {
		return ""
	}
	// get the pivot
	pivot = i - 1

	// get the lowest suffix that is lower than the pivot
	for array[j] <= array[pivot] {
		j--
	}

	// swap the pivot with the lowest suffix
	array[pivot], array[j] = array[j], array[pivot]

	// reverse the suffix to sort it into ascending order
	reversed := []rune{}
	for k := len(s) - 1; k >= i; k-- {
		reversed = append(reversed, array[k])
	}
	outputSlice := []rune{}
	outputSlice = append(array[:i], reversed...)

	return string(outputSlice)
}

func main() {
	fmt.Println(biggerIsGreater("dkhc"))
}
