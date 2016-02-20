package handler

import (
    "encoding/json"
    "github.com/jjosephy/go/sdp/httperror"
    "github.com/jjosephy/go/sdp/model"
    "github.com/jjosephy/go/sdp/converter"
    "net/http"
    "strconv"
)

func DeployHandler() http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
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

      var m model.DeployModel
      // Switch on Request Method
      switch r.Method {
          case "POST":
             switch version {
                  case 1.0:
                     json.NewEncoder(w).Encode(converter.ConvertDeployModelToContractV1(m))
                  default:
                      httperror.UnsupportedVersion(w)
                      return;
              }
          default:
              w.WriteHeader(http.StatusMethodNotAllowed)
              return
      }
    }
}
