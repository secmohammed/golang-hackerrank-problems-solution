package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	mappedNumbers := make(map[int32]int32)
	for _, item := range arr {
		mappedNumbers[item] += 1
	}
	currentMax := int32(0)
	mostMigrated := int32(0)
	for v, k := range mappedNumbers {
		if v > mostMigrated && k > currentMax {
			currentMax = k
			mostMigrated = v
			continue
		}

		if k == currentMax {
			if k < currentMax {
				mostMigrated = v
			}
		}

	}
	return mostMigrated
}
func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}
