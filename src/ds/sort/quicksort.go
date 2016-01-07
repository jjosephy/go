package sort

import (
)

type QuickSort struct {
    List []int
}

func (q *QuickSort) Sort() {
    q.quicksort(0, len(q.List) - 1)
}

func (q* QuickSort) quicksort(l int, h int) {
    if l < h {
        p := q.partition(l, h)
        q.quicksort(l, p - 1)
        q.quicksort(p + 1, h)
    }
}

func (q *QuickSort) partition(l int, r int) (int) {
    p := q.List[r]
    i := l

    for j := l; j < r; j++ {
        if q.List[j] <= p {
            t := q.List[i]
            q.List[i] = q.List[j]
            q.List[j] = t
            i++
        }
    }

    t := q.List[i]
    q.List[i] = q.List[r]
    q.List[r] = t


    return i
}
