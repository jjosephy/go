package directedGraph

import (
    "container/list"
    //"errors"
)

/*
type ErrorModel struct {
    Message string
    ErrorCode    int
}

*/

type DirectedGraph struct {
    graph []list.List
}

func (g *DirectedGraph) Initialize(size int) {
    g.graph = nil
}
