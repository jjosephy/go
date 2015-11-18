package directedGraph

import (
    "container/list"
    "errors"
    "fmt"
)

type DirectedGraph struct {
    graph []list.List //adjency list
    //graph [][]int // adjency matrix
    Edges int
    Vertices int
}

func (g *DirectedGraph) Initialize(size int) {
    g.graph = make([]list.List, size, size)
    g.Vertices = size
}

func (g *DirectedGraph) Adjacent(x int) (list.List, error) {
    var l list.List
    if x < 0 || x >= g.Vertices {
        return l, errors.New("Invalid Node")
    }

    return g.graph[x], nil
}

func (g *DirectedGraph) AddEdge(x int, y int) (error){
    if x >= g.Vertices || x < 0 || y >= g.Vertices || y < 0 {
        return errors.New("invalid edge")
    }

    g.graph[x].PushBack(y)
    g.Edges++

    return nil
}

func (g *DirectedGraph) DFS() {
    marked := make([]bool, g.Vertices, g.Vertices)
    g.dfs_int(marked, 0)
}

func (g *DirectedGraph) dfs_int(marked []bool, x int) {
    n, err := g.Adjacent(x)

    if err != nil {
        fmt.Print("invalid node")
        return
    }

    fmt.Printf("%v\n", n)
}
