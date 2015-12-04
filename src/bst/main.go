package main

import (
    "error"
    "fmt"
)
type bstcompare interface {
    Compare(x *bstnode, y *bstnode) (int, error)
}

func (n *bstnode) Compare(x *bstnode, y *bstnode) (int, error) {
    // 1 = x > y
    // 0 x == y
    // -1 x < y
}

type bstnode struct {
    data interface {}
    left *bstnode
    right *bstnode
}

func insert(node bstnode) (*bstnode, error) {


}

// Main entry point used to set up routes //
func main() {

    root := make(bstnode{ data: 5 })

    fmt.Println("main..")
}
