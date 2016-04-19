package main

import (
    "errors"
    "fmt"
)

type Heap struct {
    sequence []interface{}
    Length int
}

func (h *Heap) PrintHeap() {
    for i := 0; i < h.Length; i++ {
		fmt.Println(h.sequence[i])
	}
}

func (h *Heap) PrintLength() {
    fmt.Println(h.Length)
}

func (h *Heap) Insert(i interface{}) (error) {
    if i == nil {
        return errors.New("invalid entry")
    }

    h.sequence[h.Length] = i
    h.Length = h.Length + 1
    h.heapUp()
    return nil
}

func (h *Heap) heapUp() {
    k := h.Length
    for k > 0 {
        x := k/2
        if h.sequence[x] < h.sequence[k] {
            t := h.sequence[k]
            h.sequence[k] = h.sequence[x]
            h.sequence[x] = t
        }
        k = x
    }
}

func CreateHeap(size int) (h *Heap, e error) {
    if size < 0 {
        return nil, errors.New("invalid size")
    }

    r := &Heap {
        sequence : make([]interface{}, size),
        Length : 0,
    }

    return r, nil
}

////////----------------/////////
func main() {
    fmt.Println("main")

    heap, err := CreateHeap(10)

    if err != nil {
        fmt.Println("error")
    }

    heap.Insert(1)
    heap.Insert(2)

    heap.PrintLength()
    heap.PrintHeap()



}
