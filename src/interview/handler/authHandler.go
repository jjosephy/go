package handler

import (
    "net/http"
)

func AuthHandler(signingKey []byte) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
       w.Write([]byte("1234567890"))
  }
}
