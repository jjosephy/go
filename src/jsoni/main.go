package main

import (
    "jsond"
    "fmt"
    "reflect"
)

/*
func Stringify(m map[string]interface{}) (string) {
    s, _ := json.Marshal(m)
    return string(s)
}
*/

func PrintType(i interface{}) {
    fmt.Printf("p : %#v -- t %T", i, i)
    fmt.Println("")
}

func plm(str string) {
    fmt.Println(str)
}

func pl() {
    fmt.Println("")
}

func pls(i interface{}) {
    fmt.Printf("String: %s", i)
}

func plv(i interface{}) {
    fmt.Printf("Object: %v", i)
    pl()
}

func plk(i interface{}) {
    if i == nil {
        fmt.Printf("Kind: Is Nil")
        return
    }

    t := reflect.TypeOf(i)
    k := t.Kind()
    fmt.Printf("Kind: %s", k.String())
    pl()
}

func plt(i interface{}, s ...string) {
    if s != nil && len(s) > 0 {
        fmt.Printf("%s Type: %T", s[0], i)
    } else {
        fmt.Printf("Type: %T", i)
    }
    pl()
}
///////////////////////////////////

/////////////////////
func plb(i interface{}) {
    plv(i)
    plt(i)
    plk(i)
    pl()
}

func pln(n jsond.JsonNode) {
    pl()
    plm("-- Node --")
    plv(n)
    plt(n)
    plk(n)
    plm("-- Value --")
    plv(n.Value())
    plt(n.Value())
    plk(n.Value())
    pl()


}
// Main entry point used to set up routes //
func main() {
    //b := `{"candidate":"Candidate Name","comments":[{"content":"db Content","interviewer":"interviewer 0","interviewerId":"0"},{"content":"db Content","interviewer":"interviewer 1","interviewerId":"1"},{"content":"db Content","interviewer":"interviewer 2","interviewerId":"2"}],"complete":false,"id": 2}`
    b := `{"strings":["stringers","hokkcers","meaosfd"]}`
    root, _ := jsond.Parse(b)
    //plb(root)

    //tx := root.Property("candidate")
    //pln(tx)

    s := root.Property("strings")
    //pln(s)

    if iter := s.Iterator(); iter != nil {

        jn := iter.Last()
        pln(*jn)
        iter.First()

        for ptr := iter.Next(); iter.End != true; ptr = iter.Next() {
            //c := ptr.Property("interviewer")
            //pln(bv)
            pln(*ptr)
        }
    } else {
        fmt.Println("Iterator is nil")
    }






    //plb(iter)

    /*
    for ; sum < 1000; {
    		sum += sum
    	}
        */





    /*
    for _, o := range s.Value {
        plt(o)
    }
    */
    //l := s
    /*
    sx := reflect.ValueOf(s.Value())
    for i := 0; i < sx.Len(); i++ {
        x := sx.Index(i)
        plv(x)
        plt(x)
    }
    */

    //str := s.Value()
    //plt(str)

    //var d map[string]interface{}
    //json.Unmarshal([]byte(b), &d)

    //bt := d["comments"]

    //g := len(bt)
    //plt(bt)
    //plv(bt)
    //plt(bt, "Root")
    //plk(bt)
    //pl()

    /*
    s := reflect.ValueOf(bt)
    ix := make([]interface{}, s.Len(), s.Len())
    for i := 0; i < s.Len(); i++ {
        x := s.Index(i)
        plv(x)
        plt(x)
        ix[i] = x
    }
    pl()
    pl()

    for i := 0; i < len(ix); i++ {
        y := ix[i]
        plv(y)
        plt(y)
    }
    */

    //pl()

    //plt(ix)
    //plv(ix)
    //plt(ix, "Copied Slice")

    //plm("Looping copy")
    //for t := 0; t < len(ix); t++ {
    //    plt(ix[t])
    //}

    //col := map["comments"]
    /*
    d := Parse(b)
    p := d.m["comments"]
    */
    //fmt.Printf("p : %v", p)
    //fmt.Println("")


    //t := reflect.TypeOf(p)
    //fmt.Printf("v : %s", t.Kind())

    fmt.Println("")


    //i := len(p)
    //fmt.Printf("%s", i)

    /*
    str := Stringify(d)

    fmt.Println(str)
    */
}
