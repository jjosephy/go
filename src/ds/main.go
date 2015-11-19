package main

import (
    "graph"
)

func main() {

    g := DirectedGraph {}

    g.Initialize(5)
    g.AddEdge(0, 1)
    g.AddEdge(1, 2)

    g.DFS()
}