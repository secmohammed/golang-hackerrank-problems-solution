package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func bonAppetit(bill []int64, k, b int64) {
	if k != -1 && b != -1 {

		sum := int64(0)
		for i, v := range bill {
			if i == int(k) {
				continue
			}
			sum += v
		}
		if b == sum/2 {
			fmt.Println("Bon Appetit")
		} else {
			fmt.Println(b - sum/2)
		}
	}
}
func main() {
	var k, b int64
	var bill []int64

	flag.Int64Var(&k, "excluded", -1, "used to exclude item from the list")
	flag.Int64Var(&b, "total", -1, "used to calculate the amount that anna contributed to the bill")

	flag.Func("bill", "bill int list", func(val string) error {
		numbersString := strings.Split(val, " ")
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Print("Failed to parse candle number")
				return err
			}
			bill = append(bill, int64(number))

		}
		return nil
	})
	flag.Parse()
	bonAppetit(bill, k, b)
}
