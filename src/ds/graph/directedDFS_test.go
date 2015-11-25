package graph

import (
    "testing"
)


func Test_DirectedDFS_Success_SimpleTraverse(t *testing.T) {
    var g *DirectedGraph
    var dfs * DirectedDFS
    var e error

    if g, e = NewDirectedGraph(10); e != nil {
        t.Error("Failed to create a new graph")
        t.FailNow()
    }

    if dfs, e = NewDirectedDFS(g); e != nil {
        t.Error("Failed to create a new Directed DFS")
        t.FailNow()
    }

    g.AddEdge(0, 1)
    g.AddEdge(0, 2)
    g.AddEdge(1, 2)
    g.AddEdge(2, 3)
    g.AddEdge(3, 4)

    dfs.DFS()
}
