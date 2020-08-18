package main

import (
	"fmt"
)

func happyLadybugs(b string) bool {
	r := []rune("_")[0]
	m := map[rune]int{}
	for _, v := range b {
		m[v]++
	}
	for k := range m {
		if k != r && m[k] == 1 {
			return false
		}
	}

	if m[r] == 0 {
		for i := 1; i < len(b)-1; i++ {
			p := []rune(string(b[i-1]))[0]
			n := []rune(string(b[i+1]))[0]
			v := []rune(string(b[i]))[0]
			if p != v && n != v {
				return false
			}
		}
	}
	return true
}

func main() {
	g := "RBY_YBR"
	fmt.Println(happyLadybugs(g))
}
