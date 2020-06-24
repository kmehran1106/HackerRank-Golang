package main

import (
	"fmt"
)

func savePrisoner(prisoners, sweets, chair int) (output int) {
	if sweets > prisoners {
		sweets = sweets % prisoners
		if sweets == 0 {
			sweets = prisoners
		}
	}
	chair--
	diff := prisoners - chair
	if diff >= sweets {
		output = chair + sweets
	} else {
		output = sweets - diff
	}
	return output
}

func main() {
	fmt.Println(savePrisoner(7, 19, 2))
	fmt.Println(savePrisoner(13, 140874526, 1))
}
