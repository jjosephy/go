package contract

import (
    "testing"
)

func TestCreateContract(t *testing.T) {
    c := new(CommentContractV1)

    if c == nil {
        t.Fatal("Failed creating instance")
    }
}
