package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func gradingStudents(grades []int32) []int32 {
	var result []int32
	for _, grade := range grades {
		if grade < 38 {
			result = append(result, grade)
			continue
		}
		diff := 5 - grade%5
		if diff < 3 {
			result = append(result, grade+diff)
			continue
		}
		result = append(result, grade)
	}
	return result
}
func main() {
	var numbers []int32
	flag.Func("numbers", "int list", func(val string) error {
		numbersString := strings.Split(val, " ")
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Print("Failed to parse number")
				return err
			}
			numbers = append(numbers, int32(number))

		}
		return nil
	})
	flag.Parse()
	fmt.Println(gradingStudents(numbers))
}
