package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "interview/contract/v1"
    "interview/model"
    "strconv"
)

func resourceHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}

func writeBadRequestResponse(w http.ResponseWriter, errCode int, msg string) {
    w.WriteHeader(http.StatusBadRequest);
    json.NewEncoder(w).Encode(contract.ErrorDetail { Code: errCode, Message: msg })
}

func interviewHandler(w http.ResponseWriter, r *http.Request) {
    var version float64
    h := r.Header.Get("api-version")
    if h == "" {
        writeBadRequestResponse(w, 1001, "No Version Provided")
        return;
    } else {
        v, err := strconv.ParseFloat(h, 64)
        if err != nil {
            writeBadRequestResponse(w, 1002, "Invalid Version Provided")
            return
        }
        version = v
    }

    // Switch on Request Method
    switch r.Method {
        case "GET":
            // Get Id or Name for Search
            qId := r.URL.Query()["id"]
            qName := r.URL.Query()["cname"]

            if (qId == nil && qName == nil) {
                writeBadRequestResponse(w, 1000, "No Parameters Provided")
                return;
            }

            var id string
            if len(qId) > 0 {
                id =  qId[0]
            } else {
                id = "noId"
            }

            /*
            var name string
            if len(qName) > 0 {
                name = qName[0]
            } else {
                name = "noName"
            }
            */

            m := model.InterviewModel {
                Id: id,
            }

            switch version {
                case 1.0:
                    m.Id = "version 1"
            }

            //Find by id or name
            interview := contract.InterviewContractV1 {
                Id: m.Id,
                Interviewer: "rebuild",
                Candidate: "Bob",
            }

            json.NewEncoder(w).Encode(interview)
        case "POST":
            fmt.Fprintf(w, "POST Success")
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
    }
}

// Main entry point used to set up routes //
func main() {
    mux := http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir("../src/interview/web/")))
    mux.HandleFunc("/interview", interviewHandler)
    fmt.Println("Server Running")
    http.ListenAndServe(":8080", mux)
}
