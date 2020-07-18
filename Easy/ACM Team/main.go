package main

import (
	"fmt"
)

func getACMTeam(array []string) (output []int) {
	n := len(array)    // number of people
	m := len(array[0]) // number of topics

	maxTopics, numTeams := 0, 0

	// convert every string inside the array to a boolean array
	// 11001 -> []bool{true, true, false, false, true}
	b := [][]bool{}
	for _, v := range array {
		t := []bool{}
		for _, c := range v {
			var x bool
			if string(c) == "1" {
				x = true
			}
			t = append(t, x)
		}
		b = append(b, t)
	}

	// do the actual processing now
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			f := 0
			for k := 0; k < m; k++ {
				if b[i][k] || b[j][k] {
					f++
				}
			}
			if f == maxTopics {
				numTeams++
			} else if f > maxTopics {
				numTeams = 1
				maxTopics = f
			}
		}
	}
	output = []int{maxTopics, numTeams}
	return output
}

func main() {
	fmt.Println(getACMTeam([]string{"11101", "10101", "11001", "10111", "10000", "01110"}))
}
