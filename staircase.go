package main

import (
	"flag"
	"fmt"
)

func staircase(n int32) {
	//for i := int32(0); i < n; i++ {
	//	fmt.Println(strings.Repeat(" ", int(n-i-1)), strings.Repeat("#", int(i+1)))
	//}
	for i := int32(0); i < n; i++ {
		row := ""
		for j := int32(0); j < n-i-1; j++ {
			row = fmt.Sprintf("%s ", row)
		}
		for j := int32(0); j <= i; j++ {
			row = fmt.Sprintf("%s#", row)
		}
		fmt.Println(row)
	}
}
func main() {
	var steps int

	flag.IntVar(&steps, "step", 0, "Step count")
	flag.Parse()
	staircase(int32(steps))
}
