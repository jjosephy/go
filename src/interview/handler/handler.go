package handler

import (
    "encoding/json"
    "fmt"
    "interview/contract/v1"
    "interview/httperror"
    "interview/model"
    "interview/repository"
    "net/http"
    "strconv"
)

func InterviewHandler(data repository.InterviewRepository) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
      //TODO: validate that the passed in repository is not null

      var version float64
      h := r.Header.Get("api-version")
      if h == "" {
          httperror.NoVersionProvided(w)
          return;
      } else {
          v, err := strconv.ParseFloat(h, 64)
          if err != nil {
              httperror.InvalidVersionProvided(w)
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

              //TODO: scrub input
              if (qId == nil && qName == nil) {
                  httperror.NoQueryParametersProvided(w)
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

              // TODO: integrate with repository to build model and translate
              st := data.GetInterview(id, name)

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
                      httperror.UnsupportedVersion(w)
                      return;
              }

              // TODO: model->contract contract->model
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
}
