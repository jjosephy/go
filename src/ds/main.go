package main

import (
    //"ds/graph"
    "ds/sort"
    "fmt"
)

func main() {
    a := []int{ 4, 7, 2, 1, 0 }

    q := sort.QuickSort {
        List : a,
    }

    q.Sort()

    for i := 0; i < len(q.List); i++ {
        fmt.Print(" ", q.List[i], " ")
    }

    fmt.Println("")

    // fill out a matrix that is len of array + 1 as X and 0 to SUM as Y
    // input

    /*
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

    *//*
    func main() {
        a := []int{ 4, 7, 2, 1, 0 }
        quicksort(a, 0, 4)

        for i := 0; i < len(a); i++ {
            fmt.Print(" ", a[i], " ")
        }

        fmt.Println("")
    }
    */
}
