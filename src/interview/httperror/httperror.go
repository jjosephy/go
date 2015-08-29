package httperror

import (
    "encoding/json"
    "fmt"
    "interview/model"
    "net/http"
)

const NO_VERSION_PROVIDED               = "No Version Provided"
const INVALID_VERSION                   = "Invalid Version Provided"
const NO_PARAMETERS_PROVIDED            = "No Parameters Provided"
const UNSUPPORTED_VERSION               = "Unsupported Version Provided"
const GET_INTERVIEW_ERROR_MSG           = "Error occurred tyring to Get interview details"

// 400 Errors
const BADREQUEST_NOINPUTPARAMETERS      = 4000
const BADREQUEST_NOVERSION              = 4001
const BADREQUEST_INVALIDVERSION         = 4002
const BADREQUEST_UNSUPPORTEDVERSION     = 4003

// 500 Errors
const SERVER_ERROR_GENERAL              = 5000
const GET_INTERVIEW_FAILURE             = 5001


func GetInterviewFailed(w http.ResponseWriter, e error) {
    s := fmt.Sprint(GET_INTERVIEW_ERROR_MSG, " : ", e)
    w.WriteHeader(http.StatusInternalServerError);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: GET_INTERVIEW_FAILURE, Message: s })
}

func NoVersionProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_NOVERSION, NO_VERSION_PROVIDED)
}

func InvalidVersionProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_INVALIDVERSION, INVALID_VERSION)
}

func NoQueryParametersProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_NOINPUTPARAMETERS, NO_PARAMETERS_PROVIDED)
}

func UnsupportedVersion(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_UNSUPPORTEDVERSION, UNSUPPORTED_VERSION)
}

func writeBadRequest(w http.ResponseWriter, code int, msg string) {
    w.WriteHeader(http.StatusBadRequest);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: code, Message: msg })
}
