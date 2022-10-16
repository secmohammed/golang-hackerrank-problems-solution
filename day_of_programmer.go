package main

import (
	"flag"
	"fmt"
)

func dayOfProgrammer(year int32) string {
	monthYear := fmt.Sprintf("09.%d", year)
	var day string
	isLeap := func(year *int32) bool {
		if *year > 1918 {
			return *year%400 == 0 || *year%4 == 0 && *year%100 != 0
		}
		return *year%4 == 0
	}
	if year == 1918 {
		day = "26"
	} else {
		if isLeap(&year) {
			day = "12"
		} else {
			day = "13"
		}
	}
	return fmt.Sprintf("%s.%s", day, monthYear)
}
func main() {
	var year int
	flag.IntVar(&year, "year", 0, "Year to determine the programmer day")
	flag.Parse()
	fmt.Println(dayOfProgrammer(int32(year)))
}
