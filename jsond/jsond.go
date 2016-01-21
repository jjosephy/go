package jsond

import (
    "encoding/json"
    "reflect"
)

func createIterator(i *interface{}) (*JsonIterator) {
    s := reflect.ValueOf(*i)
    ix := make([]JsonNode, s.Len(), s.Len())
    for t := 0; t < s.Len(); t++ {
        x := s.Index(t)
        ix[t] = createNode(x.Interface())
    }

    return &JsonIterator {
        p: 0,
        l: len(ix),
        jn: ix,
    }
}

func getKind(i interface{})(reflect.Kind) {
    t := reflect.TypeOf(i)
    return t.Kind()
}

func createNode(i interface{}) (JsonNode) {
    if i == nil {
        return JsonNode{i, nil}
    }

    switch getKind(i)  {
        case reflect.Slice:
            return JsonNode{i, createIterator(&i)}
        default:
            return JsonNode{i, nil}
    }

}

func Parse(s string) (JsonRoot, error){
    var root JsonRoot
    var d map[string]interface{}
    if err := json.Unmarshal([]byte(s), &d); err != nil {
        return root, err
    }
    root = JsonRoot {d}
    return root, nil
}
