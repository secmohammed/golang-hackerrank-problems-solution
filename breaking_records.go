package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func breakingRecords(scores []int64) []int64 {
	min, max := int64(0), int64(0)
	var result []int64 // min max
	for k, score := range scores {
		if k == 0 {
			min, max = score, score
			result = append(result, 0, 0)
			continue
		}
		if score > max {
			max = score
			result[1] += 1
		}
		if min > score {
			min = score
			result[0] += 1
		}
	}
	return result
}
func main() {
	var scores []int64
	flag.Func("scores", "score int list", func(val string) error {
		numbersString := strings.Split(val, " ")
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Print("Failed to parse score number")
				return err
			}
			scores = append(scores, int64(number))

		}
		return nil
	})
	flag.Parse()
	fmt.Println(breakingRecords(scores))
}
