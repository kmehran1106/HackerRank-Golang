package main

import (
	"fmt"
)

// AppendDelete :
func AppendDelete(s, t string, k int) (output bool) {
	var minLen int
	var common int
	lenS := len(s)
	lenT := len(t)
	if lenS <= lenT {
		minLen = lenS
	} else {
		minLen = lenT
	}
	for i := 0; i < minLen; i++ {
		if s[i] != t[i] {
			break
		}
		common++
	}
	/*
		CASE1:
	        if len(s) < k - len(t), then we can delete all the letters,
	        keep on going until remaining(k) == len(t), then append len(t)
	    CASE2:
	        if len(s) + len(t) - 2*common > k,
	        then we cannot solve this case
	    CASE3:
	        if len(s) - common + len(t) - common is ODD, then k must be ODD as well
	        if len(s) - common + len(t) - common is EVEN, then k must be EVEN as well
	        otherwise, we cannot solve this case
	    CASE4:
	        NO
	*/

	if lenS < k-lenT {
		output = true
	} else if lenS+lenT-2*common > k {
		output = false
	} else if (lenS+lenT-2*common)%2 == k%2 {
		output = true
	} else {
		output = false
	}
	return output
}

func main() {
	fmt.Println(AppendDelete("qwerasdf", "qwerbsdf", 6))
}
