package graph

import (
    "testing"
)

func assertIsTrue(t *testing.T, condition bool, msg string) {
    if !condition {
        t.Error(msg)
    }
}

func Test_DirectedGraph_Fail_InvalidInitializer(t *testing.T) {
    _, e := NewDirectedGraph(-1)

    if e == nil {
        t.Error("No Error returned from Initialize")
    }
}

func Test_DirectedGraph_Fail_TestInvalidNode(t *testing.T) {
    //g, e := NewDirectedGraph(5)


}

func Test_DirectedGraph_Success_ValidInitializer(t *testing.T) {
    g, e := NewDirectedGraph(5)

    if e != nil {
        t.Error("Error returned from Initialize")
    }

    if e = g.AddEdge(0, 1); e != nil {
        t.Errorf("Error tyring to add edge &s", e.Error())
    }

    if e = g.AddEdge(1, 2); e != nil {
        t.Errorf("Error tyring to add edge &s", e.Error())
    }

    assertIsTrue(t, g.Edges == 2, "Number of edges dont match")
    //dfs(&g)
}
