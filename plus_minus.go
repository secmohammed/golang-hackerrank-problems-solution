package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func plusMinus(arr []int32) {
	counts := []int32{0, 0, 0}
	for _, val := range arr {
		if val > 0 {
			counts[0] += 1
		} else if val < 0 {
			counts[1] += 1
		} else {
			counts[2] += 1
		}

	}
	for _, val := range counts {
		fmt.Printf("%.6f\n", float32(val)/float32(len(arr)))
	}
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
	plusMinus(numbers)
}
