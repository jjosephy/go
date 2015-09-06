package jsond

import (
    "encoding/json"
    "testing"
)

const SIMPLE_JSON = `{"strings":["stringerOne","stringerTwo","stringerThree"]}`
type SimpleJson struct {
    N string `json:"name"`
    S []string `json:"strings"`
}

func AssertIsTrue(t *testing.T, c bool, msg string, e error) {
    t.Logf("c: %v", c)
    if c != true {
        if e != nil {
            t.Fatalf("Error Msg: %s - Details: %v", msg, e)
        } else {
            t.Fatalf("Error Msg: %s", msg)
        }
    }
}

func Test_Success_IterateSimpleCollection(t *testing.T) {
    sa := []string{"a", "b", "c", "d"}
    s := SimpleJson { "name", sa }
    o, e := json.Marshal(s)
    AssertIsTrue(t, e == nil, "Error trying to marshal SimpleJson", e)

    jRoot, e := Parse(string(o))
    AssertIsTrue(t, e == nil, "Error trying to parse json string", e)

    propN := jRoot.Property("strings")
    AssertIsTrue(t, propN.Value() != nil, "JsonNode has no value", nil)

    v := propN.Iterator();
    AssertIsTrue(t, v != nil, "Nil Interator on Collection", nil)
    for p := v.Next(); v.End != true; p = v.Next() {
        j := *p
        t.Logf("value: %v", j.Value())
    }
}
