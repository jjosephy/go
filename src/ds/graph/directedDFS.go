package graph

import (
    "errors"
    "fmt"
    "log"
    //"os"
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
            logMessage(fmt.Sprint("X:", x, " Y: ", e.Value.(int)))
            dfs.dfsInternal(e.Value.(int))
        }
    }
}

func logMessage(msg string) {
    //f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    //if err != nil {
    //    return
    //}
    //defer f.Close()
    //log.SetOutput(f)
    log.Println(msg)
}
