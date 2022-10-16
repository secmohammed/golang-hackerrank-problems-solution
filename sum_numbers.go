package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Println("You should insert a number to start summing separated by comma")
		os.Exit(0)
	}
	numbersString := strings.Split(os.Args[1], ",")
	if len(numbersString) == 0 {
		log.Println("Invalid Argument!")
		os.Exit(1)
	}
	initial := int64(0)
	for _, numberString := range numbersString {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Println("attempt converting number failed!")
			os.Exit(1)
		}
		initial += int64(number)

	}
	fmt.Println(initial)
}
