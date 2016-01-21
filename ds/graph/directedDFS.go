package graph

import (
    "errors"
)

type DirectedDFS struct {
    graph *DirectedGraph
    Marked []bool
}

func NewDirectedDFS(g *DirectedGraph) (*DirectedDFS, error) {
    var dfs *DirectedDFS

    if g == nil{
        return dfs, errors.New("invalid graph")
    }

    dfs = &DirectedDFS {
        graph : g,
        Marked : make([]bool, g.Vertices, g.Vertices),
    }

    return dfs, nil
}

func (dfs *DirectedDFS) DFS() {
    dfs.dfsInternal(0)
}

func (dfs *DirectedDFS) dfsInternal(x int) {
    if dfs.Marked[x] == false {
        dfs.Marked[x] = true
        n, err := dfs.graph.Adjacent(x)
        if err != nil {
            return
        }
        for e := n.Front(); e != nil; e = e.Next() {
            dfs.dfsInternal(e.Value.(int))
        }
    }
}
