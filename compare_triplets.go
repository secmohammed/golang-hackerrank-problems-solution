package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func compareTriplets(a, b []int64) []int64 {
	result := []int64{0, 0} // Alice, Bob
	if len(a) != 3 || len(b) != 3 {
		log.Fatal("Invalid Input!")
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			result[0] += 1
			continue
		}
		if a[i] < b[i] {
			result[1] += 1
			continue
		}
		if a[i] == b[i] {
			continue
		}
	}
	return result
}
func main() {
	var a, b []int64
	flag.Func("a", "score int list", func(val string) error {
		numbersString := strings.Split(val, " ")
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Print("Failed to parse score number")
				return err
			}
			a = append(a, int64(number))

		}
		return nil
	})
	flag.Func("b", "score int list", func(val string) error {
		numbersString := strings.Split(val, " ")
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Print("Failed to parse score number")
				return err
			}
			b = append(b, int64(number))

		}
		return nil
	})
	flag.Parse()
	fmt.Println(compareTriplets(a, b))
}
