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

func NewDirectedGraph(size int) (*DirectedGraph, error) {
    var g *DirectedGraph

    if size < 0 {
        return g, errors.New("invalid size")
    }

    g = &DirectedGraph {
        graph : make([]list.List, size, size),
        Vertices: size,
        Edges: 0,
    }

    return g, nil
}

func (g *DirectedGraph) Adjacent(x int) (list.List, error) {
    var l list.List

    if v, e := g.isValidNode(x); !v && e != nil {
        return l, e
    }

    return g.graph[x], nil
}

func (g *DirectedGraph) AddEdge(x int, y int) (error){
    var v bool
    var e error

    if v, e = g.isValidNode(x); !v && e != nil {
        return errors.New("x node is invalid")
    }

    if v, e = g.isValidNode(y); !v && e != nil {
        return errors.New("y node is invalid")
    }

    g.graph[x].PushBack(y)
    g.Edges++

    return nil
}

func (g *DirectedGraph) isValidNode(x int) (bool, error) {
    if x >= g.Vertices || x < 0 {
        return false, errors.New("invalid node")
    }

    return true, nil
}
