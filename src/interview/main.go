package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "interview/contract/v1"
    "interview/model"
    "interview/repository"
    "strconv"
)

const BADREQUEST_NOINPUTPARAMETERS      = 1000
const BADREQUEST_NOVERSION              = 1001
const BADREQUEST_INVALIDVERSION         = 1002
const BADREQUEST_UNSUPPORTEDVERSION     = 1003

const SERVERERROR_CANTCREATEREPOSITORY  = 2000

func resourceHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}

func writeBadRequestResponse(w http.ResponseWriter, errCode int, msg string) {
    w.WriteHeader(http.StatusBadRequest);
    json.NewEncoder(w).Encode(contract.ErrorDetailV1 { Code: errCode, Message: msg })
}

func interviewHandler(w http.ResponseWriter, r *http.Request) {
    var version float64
    h := r.Header.Get("api-version")
    if h == "" {
        writeBadRequestResponse(w, BADREQUEST_NOVERSION, "No Version Provided")
        return;
    } else {
        v, err := strconv.ParseFloat(h, 64)
        if err != nil {
            writeBadRequestResponse(w, BADREQUEST_INVALIDVERSION, "Invalid Version Provided")
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
                writeBadRequestResponse(w, BADREQUEST_NOINPUTPARAMETERS, "No Parameters Provided")
                return;
            }

            var id string
            if len(qId) > 0 {
                id =  qId[0]
            } else {
                id = ""
            }

            var name string
            if len(qName) > 0 {
                name = qName[0]
            } else {
                name = ""
            }


            //ri := repository.InterviewRepository{}
            rp, err := repository.GetRepository(repository.MOCK)

            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(contract.ErrorDetailV1 { Code: SERVERERROR_CANTCREATEREPOSITORY, Message: "Could not create Repository" })
                return
            }

            st := rp.GetInterview(id, name)

            comments := model.Comments {
                model.CommentModel { Content: st, Interviewer: "interviewer 0", InterviewerId: 0 },
                model.CommentModel { Content: st, Interviewer: "interviewer 1", InterviewerId: 1 },
                model.CommentModel { Content: st, Interviewer: "interviewer 2", InterviewerId: 2 },
            }

            // Get a model and translate that
            m := model.InterviewModel {
                Candidate: "Candidate Name",
                Id: "hardcodedid",
                Comments: comments,
            }

            switch version {
                case 1.0:
                    m.Id = "version 1"
                default:
                    writeBadRequestResponse(w, BADREQUEST_UNSUPPORTEDVERSION, "Unsupported Version")
            }

            //Find by id or name
            json.NewEncoder(w).Encode(contract.InterviewContractV1 {
                Id: m.Id,
                Interviewer: "rebuild",
                Candidate: "Bob",
                Comments: contract.CommentsV1 {
                    contract.CommentContractV1  { Content: "Content", Interviewer: "interviewer 0", InterviewerId: "0" },
                    contract.CommentContractV1  { Content: "Content", Interviewer: "interviewer 1", InterviewerId: "1" },
                    contract.CommentContractV1  { Content: "Content", Interviewer: "interviewer 2", InterviewerId: "2" },
                },
            })


        case "POST":
            fmt.Fprintf(w, "POST Success")
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
    }
}

func createServer() {
    mux := http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir("../src/interview/web/")))
    mux.HandleFunc("/interview", interviewHandler)
    fmt.Println("Server Running")
    http.ListenAndServe(":8080", mux)
}

// Main entry point used to set up routes //
func main() {
    createServer()
}
