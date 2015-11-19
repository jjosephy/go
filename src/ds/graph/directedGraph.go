package graph

import (
    "container/list"
    "errors"
)

type DirectedGraph struct {
    graph []list.List //adjency list
    //graph [][]int // adjency matrix
    Edges int
    Vertices int
}

func (g *DirectedGraph) Initialize(size int) (error) {
    if size < 0 {
        return errors.New("invalid size")
    }

    g.graph = make([]list.List, size, size)
    g.Vertices = size

    return nil
}

func (g *DirectedGraph) Adjacent(x int) (list.List, error) {
    var l list.List

    if e := g.validateNode(x); e != nil {
        return l, e
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

func (g *DirectedGraph) validateNode(x int) (error) {
    if x >= g.Vertices || x < 0 {
        return errors.New("invalid")
    }

    return nil
}
