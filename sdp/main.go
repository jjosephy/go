package main

import (
    "fmt"
    "github.com/jjosephy/go/sdp/handler"
    "net/http"
)

const (
    PORT       = ":8444"
    PRIV_KEY   = "./private_key"
    PUBLIC_KEY = "./cert.pem"
)

// Main entry point used to set up routes //
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/deploy", handler.DeployHandler())
    err := http.ListenAndServeTLS(PORT, PUBLIC_KEY, PRIV_KEY, mux)
    if err != nil {
       fmt.Printf("main(): %s\n", err)
    }
}
