package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

func calculateChildAgeThroughHighestCandles(candles []int32) int32 {
	count, max := int32(0), int32(0)

	for _, candle := range candles {
		if candle > max {
			max = candle
			count = 0
		}
		if max == candle {
			count++
		}
	}
	return count

}
func main() {
	var candles []int32
	flag.Func("candles", "candle int list", func(val string) error {
		numbersString := strings.Split(val, " ")
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Print("Failed to parse candle number")
				return err
			}
			candles = append(candles, int32(number))

		}
		return nil
	})
	flag.Parse()
	calculateChildAgeThroughHighestCandles(candles)

}
