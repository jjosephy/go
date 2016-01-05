package main

import (
    //"ds/graph"
    "ds/dp"
    "fmt"
)

func main() {
    // fill out a matrix that is len of array + 1 as X and 0 to SUM as Y
    // input
    a := []int { 4, 11, 7, 9, 1, 3 }
    s := 14

    sub, _ := dp.NewSubSetSum(a, s)
    b := sub.IsSubSetSum()

    //Show Result
    sub.Dump()
    fmt.Println("")
    l := sub.List
    if b {
        fmt.Println("Solution found")
        for i := 0; i < len(l); i++ {
            fmt.Print(l[i])
            if i < len(l) - 1 {
                fmt.Print(" -> ")
            }
        }
    } else {
        fmt.Println("Solution not found")
    }
    fmt.Println("")
    fmt.Println("")
}
