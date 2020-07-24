package main

import (
	"fmt"
	"strconv"
	"strings"
)

func modifiedKaprekarNumber(p, q int) (output string) {
	for num := p; num <= q; num++ {
		d := len(strconv.Itoa(num))
		numSquare := strconv.Itoa(num * num)
		r := numSquare[len(numSquare)-d:]
		l := strings.Replace(numSquare, r, "", 1)
		if l == "" {
			l = "0"
		}
		intR, _ := strconv.Atoi(r)
		intL, _ := strconv.Atoi(l)
		if intR+intL == num {
			output += strconv.Itoa(num) + " "
		}
	}
	output = strings.TrimSpace(output)
	return output
}

func main() {
	fmt.Println(modifiedKaprekarNumber(1, 100))
}
