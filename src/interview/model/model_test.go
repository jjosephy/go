package model

import (
    "testing"
)

func Test_Success_CreateErrorModel(t *testing.T) {
    m := ErrorModel {
        Message: "ErrMessage",
        ErrorCode: 1000,
    }

    if m.Message != "ErrMessage" {
        t.Fatal("failed")
    }
}
