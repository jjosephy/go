package dp

import (
    "testing"
)

const SIZE = 25
var fib = []int { 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, }
func Test_Simple_Sequence(t *testing.T) {
    f := FibonacciSequence {}
    f.CreateSequence(SIZE)

    if f.Length != SIZE {
        t.Error("Size does not match")
    }

    for x := 0; x < SIZE; x++ {
        if f.Sequence[x] != fib[x] {
            t.Error("No match")
        }
    }
}
