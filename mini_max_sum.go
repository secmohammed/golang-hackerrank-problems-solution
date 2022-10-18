package main

import (
    "fmt"
    "math"
)

func miniMaxSum(arr []int32) {
    //result := []int32{0, 0}
    //sort.Slice(arr, func(i, j int) bool {
    //	return arr[i] < arr[j]
    //})
    //lowerMaxInts := arr[:len(arr)-1]
    //higherMaxInts := arr[1:]
    //for _, v := range lowerMaxInts {
    //	result[0] += v
    //}
    //for _, v := range higherMaxInts {
    //	result[1] += v
    //}
    //
    //return result
    sum := int64(0)
    min := int64(math.MaxInt64)
    max := int64(0)

    for _, val := range arr {
        sum += int64(val)
        if int64(val) < min {
            min = int64(val)
        }
        if int64(val) > max {
            max = int64(val)
        }
    }
    fmt.Printf("%d %d\n", sum-max, sum-min)
}

func main() {
    miniMaxSum([]int32{1, 2, 3, 4, 5})

}
