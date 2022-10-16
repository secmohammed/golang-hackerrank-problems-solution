package main

import (
	"fmt"
	"math"
)

func calculateDiagonalDifference(arr [][]int32) int32 {
	ltrd := int32(0) // left to right diagonal
	rtld := int32(0) // right to left diagonal

	// arr[0][0]
	// arr[1][1]
	// arr[2][2]
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > i {
			ltrd += arr[i][i]
			rtld += arr[i][len(arr[i])-i-1]
		}
	}

	return int32(math.Abs(float64(ltrd) - float64(rtld)))
}

func main() {
	fmt.Println(calculateDiagonalDifference([][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}))
}
