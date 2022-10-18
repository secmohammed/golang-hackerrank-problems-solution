package main

import (
    "flag"
    "fmt"
    "log"
)

func kangaroo(x1, v1, x2, v2 int64) string {

    if x1 > x2 {
        return findResult(x1, v1, x2, v2)
    }
    return findResult(x2, v2, x1, v1)
}
func findResult(x1, v1, x2, v2 int64) string {
    for x1 > x2 {
        x1 += v1
        x2 += v2
        if x1 == x2 {
            return "YES"
        }
    }
    return "NO"
}
func main() {
    var x1, v1, x2, v2 int64
    flag.Int64Var(&x1, "x1", 0, "")
    flag.Int64Var(&v1, "v1", 0, "")
    flag.Int64Var(&x2, "x1", 0, "")
    flag.Int64Var(&v2, "v2", 0, "")
    if x1 == 0 || v1 == 0 || v2 == 0 || x2 == 0 {
        log.Fatal("All of the passed arguments should be more than 0")
    }
    fmt.Println(kangaroo(x1, v1, x2, v2))
}
