package main

import (
    "flag"
    "fmt"
)

func countingValleys(steps int32, path string) int32 {
    count, level := int32(0), int32(0)
    for i := int32(0); i < steps; i++ {

        if path[i] == 85 { // current index of path equals to U
            if level == -1 {
                count++
            }
            level++
        } else {
            level--
        }
    }
    return count
}
func main() {
    var path string
    var steps int

    flag.IntVar(&steps, "step", 0, "Step count")
    flag.StringVar(&path, "path", "", "U for up and D for down.")
    flag.Parse()
    fmt.Println(countingValleys(int32(steps), path))

}
