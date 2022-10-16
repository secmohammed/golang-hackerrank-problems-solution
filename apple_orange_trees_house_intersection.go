package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// s => start of house
// t => end of house
// a => point of apple tree
// b => point of orange tree
func countApplesAndOranges(s, t, a, b int64, apples, oranges []int64) {
	appleCount := int64(0)
	for _, val := range apples {
		d := a + val
		if d >= s && d <= t {
			appleCount++
		}
	}
	fmt.Println(appleCount)
	orangeCount := int64(0)
	for _, val := range oranges {
		d := b + val
		if d >= s && d <= t {
			orangeCount++
		}
	}
	fmt.Println(orangeCount)
}
func main() {
	var s, t, a, b int64
	var apples, oranges []int64
	flag.Int64Var(&s, "start", 0, "start of the house point")
	flag.Int64Var(&t, "end", 0, "end of the house point")
	flag.Int64Var(&a, "applePoint", 0, "apple Point")
	flag.Func("apples", "apple int list", func(val string) error {
		numbersString := strings.Split(val, " ")
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Print("Failed to parse apple number")
				return err
			}
			apples = append(apples, int64(number))

		}
		return nil
	})
	flag.Func("oranges", "apple int list", func(val string) error {
		numbersString := strings.Split(val, " ")
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Print("Failed to parse apple number")
				return err
			}
			oranges = append(oranges, int64(number))

		}
		return nil
	})
	flag.Int64Var(&b, "orangePoint", 0, "orange Point")
	flag.Parse()
	if t != 0 && s != 0 && a != 0 && b != 0 && len(apples) > 0 && len(oranges) > 0 {
		countApplesAndOranges(s, t, a, b, apples, oranges)
		return
	}
	log.Fatal("Failed to calculate the count")

}
