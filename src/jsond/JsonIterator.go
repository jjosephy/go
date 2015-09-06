package jsond

import (
)

type JsonIterator struct {
    p int
    l int
    jn []JsonNode
    End bool

}

func (j *JsonIterator) Next() (*JsonNode) {
    if j == nil {
        return nil
    }

    if j.p == j.l {
        j.End = true
        return nil
    }

    v := j.jn[j.p]
    j.p++

    return &v
}

func (j *JsonIterator) Last() (*JsonNode) {
    if j == nil {
        return nil
    }

    j.p = j.l - 1
    j.End = true
    return &j.jn[j.l - 1]
}

func (j *JsonIterator) First() (*JsonNode) {
    if j == nil {
        return nil
    }

    j.p = 0
    j.End = false
    return &j.jn[0]
}
