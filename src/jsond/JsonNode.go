package jsond

import (
    //"fmt"
    "reflect"
)

type JsonNode struct {
    i interface{}
    it *JsonIterator
}

func (n *JsonNode) Iterator() (*JsonIterator) {
    return n.it
}

func (n *JsonNode) Property(str string) (JsonNode) {
    switch getKind(n.Value()) {
        case reflect.Map:
            m := n.i.(map[string]interface{})
            return createNode(m[str])
        case reflect.String:
            return createNode(n.i.(string))
        default:
            return createNode(nil)
    }
}

func (n *JsonNode) Value() (interface {}) {
    if n == nil || n.i == nil {
        return nil
    }

    return n.i
}
