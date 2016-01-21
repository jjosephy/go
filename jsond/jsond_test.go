package jsond

import (
    "encoding/json"
    "math/rand"
    "reflect"
    "testing"
)

const SIMPLE_JSON = `{"strings":["stringerOne","stringerTwo","stringerThree"]}`
type SimpleJson struct {
    N string `json:"name"`
    S []string `json:"strings"`
}

type ComplexJson struct {
    Map map[string]interface{}
    Slice [][]string
    SimpleSlice []SimpleJson
    Simple SimpleJson
    Floats []float64
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func AssertIsTrue(t *testing.T, c bool, msg string, e interface{}) {
    if c != true {
        if e != nil {
            t.Fatalf("Error Msg: %s - Details: %v", msg, e)
        } else {
            t.Fatalf("Error Msg: %s", msg)
        }
    }
}

func CreateFloatSlice() []float64 {
    f := make([]float64, 5)
    for t := range f {
        f[t] = rand.Float64()
    }

    return f
}

func CreateSimpleSlice() []SimpleJson {
    x := 5
    slice := make([]SimpleJson, x)
    for t := range slice {
        slice[t].S = []string{RandStringRunes(5), RandStringRunes(15), RandStringRunes(9), RandStringRunes(8)}
        slice[t].N = RandStringRunes(9)
    }

    return slice
}

func CreateMap() ( map[string]interface{} ){
    var m map[string] interface{}

    m = make(map[string] interface{})
    m["string"] = "zero"
    m["int"] = 0
    m["simple"] = SimpleJson {
        S: []string{ RandStringRunes(5), RandStringRunes(15), RandStringRunes(9), RandStringRunes(8) },
        N: RandStringRunes(5),
    }

    return m
}

func Test_Success_CreateComplexJsonNode(t *testing.T) {
    i := 10
    slice := make([][]string, i)

    for x := range slice {
        slice[x] = make([]string, i)
        for j := range slice[x] {
            slice[x][j] = RandStringRunes(5)
        }
    }

    sa := []string{"and", "boy", "cat", "dog"}
    c := ComplexJson {
        Slice: slice,
        Simple: SimpleJson {
            "name",
            sa,
        },
        SimpleSlice: CreateSimpleSlice(),
        Map: CreateMap(),
        Floats: CreateFloatSlice(),
    }

    o, e := json.Marshal(c)
    AssertIsTrue(t, e == nil, "Error trying to marshal SimpleJson", e)

    jRoot, e := Parse(string(o))
    AssertIsTrue(t, e == nil, "Error trying to parse json string", e)

    propN := jRoot.Property("Map")
    AssertIsTrue(t, propN.Value() != nil, "JsonNode Map has no value", nil)

    propM := propN.Property("simple")
    AssertIsTrue(t, propN.Value() != nil, "JsonNode strings has no value", nil)

    propS := propM.Property("strings")
    k := getKind(propS.Value())
    AssertIsTrue(t, k == reflect.Slice, "propS slice is not well formed", nil)
}

func Test_Success_CreateJsonNode(t *testing.T) {
    jRoot, e := Parse(SIMPLE_JSON)
    AssertIsTrue(t, e == nil, "Failed trying to parse Simple Json String", e)

    jNode := jRoot.Property("strings")
    AssertIsTrue(t, jNode.Value() != nil, "Value is Nil", jNode)
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

    p := v.Last()
    AssertIsTrue(t, p.Value() != nil, "Nil value returned by Last()", nil)
    AssertIsTrue(t, v.End == true, "End should be true", v.End)

    p = v.First()
    AssertIsTrue(t, p.Value() != nil, "Nil value returned by First()", nil)
    AssertIsTrue(t, v.End == false, "End should be false", v.End)

    AssertIsTrue(t, v != nil, "Nil Interator on Collection", nil)
    for p := v.Next(); v.End != true; p = v.Next() {
        AssertIsTrue(t, p != nil, "Next returned Nil Pointer", nil)
    }
}
