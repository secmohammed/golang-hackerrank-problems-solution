package main

import (
	"fmt"
)

func getMoneySpent(keyboards []int32, drives []int32, budget int32) int32 {
	maxSpent := int32(-1)
	for i := len(keyboards) - 1; i >= 0; i-- {
		for j := len(drives) - 1; j >= 0; j-- {
			total := keyboards[i] + drives[j]
			if total == budget {
				return total
			}
			if total < budget && total > maxSpent {
				maxSpent = total
			}
		}
	}
	return maxSpent
	//var result []int32
	//for i := 0; i <= len(keyboards)-1; i++ {
	//	if keyboards[i] >= budget {
	//		break
	//	}
	//	for j := 0; j <= len(drives)-1; j++ {
	//		if keyboards[i]+drives[j] <= budget {
	//			result = append(result, keyboards[i]+drives[j])
	//		}
	//	}
	//}
	//if len(result) > 0 {
	//	sort.Slice(result, func(i, j int) bool { return result[i] > result[j] })
	//	return result[0]
	//}
	//return -1
}
func main() {
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
}
