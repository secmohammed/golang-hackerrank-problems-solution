package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func simpleArraySum(ar []int32) int32 {
	sum := int32(0)
	for _, v := range ar {
		sum += v
	}
	return sum
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
	fmt.Println(simpleArraySum(numbers))

}
