package main

import "fmt"

/**

Given an integer array nums, find a subarray
that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.


Example 1:

Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.


Example 2:

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Input: nums = [-2,0,9, 3,-1]
Output: 27
Explanation: [9,3] has the largest product 27

Input: nums = [3 5 -2 -4]
Output: 120
Explanation: [3,5,-2,-4] has the largest product 120


*/
func getMaxProduct(sample []int) int {
	max := 0
	currentSequence := 1
	for i, v := range sample {
		if v*currentSequence > currentSequence || (v < 0 && len(sample)-1 >= i+1 && sample[i+1] < 0) {
			max = currentSequence * v
			currentSequence = max
		} else {
			currentSequence = 1

		}
	}
	return max
}
func main() {
	samples := [][]int{
		{2, 3, -2, 4},
		{-2, 0, -1},
		{-2, 0, 9, 3, -1},
		{3, 5, -2, -4},
		{1, 3, 0, 2, 9, -1, 3, 2, 5},
	}
	for _, sample := range samples {
		fmt.Println(getMaxProduct(sample))
	}
}
