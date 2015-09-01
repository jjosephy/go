package main

import (
    "fmt"
    "interview/handler"
    "interview/repository"
    "net/http"
)

// Main entry point used to set up routes //
func main() {
    mux := http.NewServeMux()

    // TODO: figure out path and a better way to configure
    mux.Handle("/", http.FileServer(http.Dir("../src/interview/web/")))

    // TODO: figure out injection pattern and config
    repo := repository.DBInterviewRepository{ Uri: "mongodb://localhost" }
    mux.HandleFunc("/interview", handler.InterviewHandler(&repo))

    fmt.Println("Server Running")
    //defer repo.DBSession.Close()

    http.ListenAndServe(":8080", mux)
}
