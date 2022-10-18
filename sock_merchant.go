package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func sockMerchant(n int32, ar []int32) int32 {
	result := make(map[int32]int)

	for _, sock := range ar {
		result[sock] += 1
	}
	fmt.Println(result)
	pairs := 0
	for _, v := range result {
		diff := v % 2
		if diff == 0 {
			fmt.Println(v)
			pairs += v / 2
			continue
		}
		fmt.Println(diff, v)
		pairs += (v - diff) / 2

	}

	return int32(pairs)
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
	fmt.Print(sockMerchant(int32(len(numbers)), numbers))
}
