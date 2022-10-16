package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func divisbleSumPairs(n int32, k int32, ar []int32) int32 {
	// n -> length of arr
	// k -> integer divisor
	// arr -> dataset
	count := int32(0)
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	for i := int32(0); i < n-1; i++ {
		for j := i + 1; j < n-1; j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}
func main() {
	var divisible int
	var numbers []int32
	flag.IntVar(&divisible, "divisible", 0, "divisible number that we gonna divide numbers by")
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
	fmt.Println(divisbleSumPairs(int32(len(numbers)), int32(divisible), numbers))
}
