package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func birthday(s []int32, d, m int32) int32 {
	count := int32(0)
	for i := 0; i < len(s)-int(m-1); i++ {
		sum := int32(0)
		for j, k := i, int32(0); k < m; j, k = j+1, k+1 {
			sum += s[j]
		}
		if sum == d {
			count++
		}

	}
	return count
}

func main() {
	var numbers []int32
	var m, d int
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
	flag.IntVar(&m, "m", 0, "")
	flag.IntVar(&d, "m", 0, "")

	flag.Parse()
	fmt.Println(birthday(numbers, int32(m), int32(d)))
}
