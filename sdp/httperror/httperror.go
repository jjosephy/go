package httperror

import (
    "encoding/json"
    "github.com/jjosephy/go/sdp/model"
    "net/http"
)

const (
    MSG_NO_VERSION_PROVIDED               = "No Version Provided"
    MSG_INVALID_VERSION                   = "Invalid Version Provided"
    MSG_NO_PARAMETERS_PROVIDED            = "No Parameters Provided"
    MSG_NOREQUESTBODY                     = "No Request Body Provided"
    MSG_UNSUPPORTED_VERSION               = "Unsupported Version Provided"
)

const (
    // 400 Errors
    BADREQUEST_NOINPUTPARAMETERS                  = 3000
    BADREQUEST_NOVERSION                          = 3001
    BADREQUEST_INVALIDVERSION                     = 3002
    BADREQUEST_NOREQUESTBODY                      = 3003
    BADREQUEST_INVALID_REQUESTBODY                = 3004
    BADREQUEST_UNSUPPORTEDVERSION                 = 3005

    // 500 Errors
    SERVERERROR_GENERAL                          = 5000
)

func UnsupportedVersion(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_UNSUPPORTEDVERSION, MSG_UNSUPPORTED_VERSION)
}

func GeneralServerError(w http.ResponseWriter, msg string) {
    w.WriteHeader(http.StatusInternalServerError);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: SERVERERROR_GENERAL, Message: msg })
}

func InvalidVersionProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_INVALIDVERSION, MSG_INVALID_VERSION)
}

func NoQueryParametersProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_NOINPUTPARAMETERS, MSG_NO_PARAMETERS_PROVIDED)
}

func NoVersionProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_NOVERSION, MSG_NO_VERSION_PROVIDED)
}

func InvalidRequestBody(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_INVALID_REQUESTBODY, MSG_NOREQUESTBODY)
}

func NoRequestBody(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_NOREQUESTBODY, MSG_NOREQUESTBODY)
}

func writeBadRequest(w http.ResponseWriter, code int, msg string) {
    w.WriteHeader(http.StatusBadRequest);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: code, Message: msg })
}
