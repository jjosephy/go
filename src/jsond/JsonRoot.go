package jsond

import (

)

type JsonRoot struct {
    m map[string]interface{}
}

func (r *JsonRoot) Property(str string) (JsonNode) {
    return createNode(r.m[str])
}
