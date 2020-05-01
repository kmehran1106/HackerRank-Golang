package time_conversion

import (
	"strings"
	"strconv"
)


func TimeConversion (s string) string {
	first := s[0:2]
	last := s[len(s)-2:]

	h, _ := strconv.Atoi(first)

	var n string
	if last == "PM" && h < 12 {
		n = strconv.Itoa(h+12)
	} else if last == "PM" && h == 12 {
		n = strconv.Itoa(h)
	} else if last == "AM" && h == 12 {
		n = "00"
	} else {
		n = "0" + strconv.Itoa(h)
	}

	r := strings.Replace(s, first, n, -1)
	r = strings.Replace(r, last, "", -1)

	return r
} 
