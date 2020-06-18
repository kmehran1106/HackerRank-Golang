package main

import (
	"fmt"
	"strconv"
)

func ProgrammersDay(year int) (date string) {
	isLeap := false
	if (
		(year < 1918 && year % 4 == 0) || 
		(year % 4 == 0 && year % 100 != 0) || 
		(year % 400 == 0)) {
		isLeap = true
	}
	
	// month := 9
	baseDays := 243
	if year == 1918 {
		baseDays = 230
	}
	countDays := 256
	if isLeap {
		countDays = 255
	}
	day := (countDays - baseDays)

	date = strconv.Itoa(day) + ".09." + strconv.Itoa(year)
	return date

}

func main() {
	fmt.Println(ProgrammersDay(2017))
}