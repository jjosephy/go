package dp

import (
    "errors"
)

type FibonacciSequence struct {
    Sequence []int
    Length int
    found []int
}

func (f *FibonacciSequence) CreateSequence(size int) (error) {
    if (size < 0) {
        return errors.New("Invalid size")
    }

    f.Sequence = make([]int, size)
    f.Length = size
    f.found = make([]int, size)

    for x := 0; x < size; x++ {
        f.Sequence[x] = f.fib(x)
    }

    return nil
}

// recursive dp
func (f *FibonacciSequence) fib(i int) (int) {
    if f.found[i] != 0 {
        return f.found[i]
    }

    t := i
    if i <= 0 {
        t = 0
    } else if i == 1 {
        t = 1
    } else {
        t = f.fib(i - 1) + f.fib(i - 2)
    }

    f.found[i] = t
    return f.found[i]
}
