package handler

import (
    "encoding/json"
    "fmt"
    "interview/converter"
    "interview/httperror"
    //"interview/model"
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

              // TODO: //Find by id or name
              model := data.GetInterview(id, name)

              switch version {
                  case 1.0:
                     ct := converter.ConvertModelToContractV1(model)
                     json.NewEncoder(w).Encode(ct)
                  default:
                      httperror.UnsupportedVersion(w)
                      return;
              }


          case "POST":
              fmt.Fprintf(w, "POST Success")
          default:
              w.WriteHeader(http.StatusMethodNotAllowed)
              return
      }
    }
}
