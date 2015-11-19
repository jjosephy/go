package graph

import (
    "testing"
)

func Test_DirectedGraph_Fail_InvalidInitializer(t *testing.T) {
    g := DirectedGraph {}
    e := g.Initialize(-1)

    if e == nil {
        t.Error("No Error returned from Initialize")
    }

    t.Log(e)
}

func Test_DirectedGraph_Success_ValidInitializer(t *testing.T) {
    g := DirectedGraph {}
    e := g.Initialize(5)

    if e != nil {
        t.Error("Error returned from Initialize")
    }
    g.AddEdge(0, 1)
    g.AddEdge(1, 2)

    //dfs(&g)
}
