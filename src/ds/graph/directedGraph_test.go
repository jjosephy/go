package directedGraph

import (
    "fmt"
    "testing"
)

func Test_DirectedGraph(t *testing.T) {
    t.Log("Testing Directed Graph")
    g := DirectedGraph {}

    g.Initialize(5)
    g.AddEdge(0, 1)
    g.AddEdge(1, 2)

    dfs(&g)
}


func dfs(g *DirectedGraph) {
    marked := make([]bool, g.Vertices, g.Vertices)
    dfs_int(g, marked, 0)
}

func dfs_int(g *DirectedGraph, marked []bool, x int) {
    n, err := g.Adjacent(x)

    if err != nil {
        fmt.Print("invalid node")
        return
    }

    fmt.Printf("%v\n", n)


}
