package main

import (
	"strings"
	"strconv"
	"bufio"
	"os"
	"fmt"
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

func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Standard 12h Time : ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(TimeConversion(s))
}