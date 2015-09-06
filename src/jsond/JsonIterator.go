package jsond

import (
)

type JsonIterator struct {
    p int
    s []interface {}
    l int
    End bool
}

func (j *JsonIterator) Next() (JsonNode) {
    if j.p == j.l {
        j.End = true
        return createNode(nil)
    }

    v := j.s[j.p]
    j.p++

    return createNode(v)
}

func (j *JsonIterator) Last() (JsonNode) {
    j.p = j.l
    j.End = true
    return createNode(j.s[j.l])
}

func (j *JsonIterator) First() (JsonNode) {
    j.p = 0
    j.End = false
    return createNode(j.s[0])
}
