package main

import (
    "fmt"
    "interview/handler"
    "interview/repository"
    "io/ioutil"
    "net/http"
)

const (
    PORT       = ":8443"
    PRIV_KEY   = "./private_key"
    PUBLIC_KEY = "./cert.pem"
)

// Main entry point used to set up routes //
func main() {
    var e error
    var signingKey []byte
    if signingKey, e = ioutil.ReadFile("cert.pem"); e != nil {
        panic(e)
    }

    mux := http.NewServeMux()
    // TODO: figure out path and a better way to configure
    mux.Handle("/", http.FileServer(http.Dir("/home/jjosephy/Source/go/src/interview/web")))
    // TODO: figure out injection pattern and config
    repo := repository.DBInterviewRepository{ Uri: "mongodb://localhost" }
    mux.HandleFunc("/interview", handler.InterviewHandler(&repo))
    mux.HandleFunc("/token", handler.TokenHandler(signingKey))
    mux.HandleFunc("/auth", handler.AuthHandler(signingKey))

    fmt.Println("Server Running")

    err := http.ListenAndServeTLS(PORT, PUBLIC_KEY, PRIV_KEY, mux)
    if err != nil {
       fmt.Printf("main(): %s\n", err)
   }
}
