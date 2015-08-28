package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "io/ioutil"
    "encoding/json"
    "interview/contract/v1"
)

func initialize() {

}

func TestBadRequest_NoVersion(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(interviewHandler))
    defer ts.Close()

    resp, err := http.Get(ts.URL)
    if err != nil {
        t.Fatalf("Unexpected error %s", err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Unexpected error reading body %s\n", err)
    }
    var errDetail contract.ErrorDetailV1
    json.Unmarshal(body, &errDetail)

    if err != nil {
        t.Fatalf("Error deserializing json %s\n", err)
    }

    if errDetail.Code != 1001 {
        t.Fatalf("Error Code is not correct expected 1001, got %s\n", errDetail.Code)
    }

    if resp.StatusCode != 400 {
        t.Fatalf("Received Unexpected Status Code: %d\n", resp.StatusCode)
    }
}
