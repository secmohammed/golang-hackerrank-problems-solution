package main

import (
    "fmt"
)

// get the least hops that we need to flip pages to reach certain page out of total number of given pages.
func pageCount(n int32, p int32) int32 {

    lowerLimit := int32(2)
    getCount := func(lowerLimit *int32, p int32) int32 {
        count := int32(0)
        for p >= *lowerLimit {
            count++
            p -= 2

        }
        return count

    }
    if p > n/2 {
        p = n - p
        if n%2 == 0 {
            lowerLimit = 1
        }
    }
    return getCount(&lowerLimit, p)
}
func main() {
    fmt.Println(pageCount(5, 4))
}
