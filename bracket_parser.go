package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:
Input: s = "()"
Output: true


Example 2:
Input: s = "()[]{}"
Output: true


Example 3:
Input: s = "(]"
Output: false

*/
func appendToPositions(positions *[][]int, currentIndex, value int) bool {
	p := *positions
	if len(p)-1 >= currentIndex && len(p[currentIndex]) == 2 {
		return false
	}
	p[currentIndex] = append(p[currentIndex], value)

	positions = &p
	return true
}
func isValid(str string) bool {
	if !strings.ContainsAny(str, "[]{}()") {
		return false
	}
	// means that we have unclosed bracket.
	if len(str)%2 == 1 {
		return false
	}
	var positions [][]int
	currentCurlyIndex, currentSquareIndex, currentParenthesisIndex := 0, 0, 0

	for i, v := range str {
		if string(v) == "(" {
			positions = append(positions, []int{})
			currentCurlyIndex = len(positions) - 1
			appendToPositions(&positions, currentCurlyIndex, i)
		} else if string(v) == ")" && !appendToPositions(&positions, currentCurlyIndex, i) {
			return false
		}
		if string(v) == "[" {
			positions = append(positions, []int{})
			currentSquareIndex = len(positions) - 1
			appendToPositions(&positions, currentSquareIndex, i)
		} else if string(v) == "]" && !appendToPositions(&positions, currentSquareIndex, i) {
			return false
		}
		if string(v) == "{" {
			positions = append(positions, []int{})
			currentParenthesisIndex = len(positions) - 1
			appendToPositions(&positions, currentParenthesisIndex, i)
		} else if string(v) == "}" && !appendToPositions(&positions, currentParenthesisIndex, i) {
			return false
		}
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i][0] < positions[j][0]
	})
	for i := 0; i < len(positions); i++ {
		for j := i; j < len(positions); j++ {
			if positions[i][1] > positions[j][1] {
				return false
			}
		}
	}
	return true
}

func main() {

	samples := []string{"[](){}", "[(])", "(][)", "([])", "1x22"}
	for _, v := range samples {
		fmt.Println(fmt.Sprintf("%s is %v", v, isValid(v)))

	}
}
