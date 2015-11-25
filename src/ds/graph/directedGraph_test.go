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
    g, e := NewDirectedGraph(5)

    if e = g.AddEdge(-1, 0); e == nil {
        t.Errorf("error should not be nil %v", e)
    }
    assertIsTrue(t, e.Error() == "x node is invalid", "incorrect error")

    if e = g.AddEdge(5, 0); e == nil {
        t.Error("error should not be nil")
    }
    assertIsTrue(t, e.Error() == "x node is invalid", "incorrect error")

    if e = g.AddEdge(0, -1); e == nil {
        t.Error("error should not be nil")
    }
    assertIsTrue(t, e.Error() == "y node is invalid", "incorrect error")

    if e = g.AddEdge(0, 5); e == nil {
        t.Error("error should not be nil")
    }
    assertIsTrue(t, e.Error() == "y node is invalid", "incorrect error")
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
}
