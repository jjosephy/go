package sort

import (
    "testing"
)


func Test_QuickSort_Success_GeneralSort(t *testing.T) {
    a := []int{ 4, 7, 2, 1, 0 }
    b := []int{ 0, 1, 2, 4, 7 }

    q := QuickSort {
        List : a,
    }

    q.Sort()

    if len(a) != len(q.List) {
        t.Error("invalid length")
    }

    for i := 0; i < len(q.List); i++ {
        if q.List[i] != b[i] {
            t.Error("Sort does not match")
        }
    }
}
