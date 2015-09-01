package httperror

import (
    "encoding/json"
    "fmt"
    "interview/model"
    "net/http"
)

const MSG_NO_VERSION_PROVIDED               = "No Version Provided"
const MSG_INVALID_VERSION                   = "Invalid Version Provided"
const MSG_NO_PARAMETERS_PROVIDED            = "No Parameters Provided"
const MSG_UNSUPPORTED_VERSION               = "Unsupported Version Provided"
const MSG_GET_INTERVIEW_ERROR               = "Error occurred tyring to Get interview details"
const MSG_SAVE_INTERVIEW_ERROR              = "Error occurred trying to Save the interview"
const MSG_FAILED_READING_BODY               = "Failed to parse the request body"
const MSG_INTERVIEW_NOTFOUND                = "Interview not found"

// 400 Errors
const BADREQUEST_NOINPUTPARAMETERS                  = 3000
const BADREQUEST_NOVERSION                          = 3001
const BADREQUEST_INVALIDVERSION                     = 3002
const BADREQUEST_UNSUPPORTEDVERSION                 = 3003
const BADREQUEST_FAILED_DECODING_REQUEST_BODY       = 3004

// 404 Errors
const NOTFOUND_INTERVIEW_NOTFOUND                   = 4000

// 500 Errors
const SERVERERROR_GENERAL                          = 5000
const SERVERERROR_GET_INTERVIEW_FAILURE            = 5001
const SERVERERROR_SAVE_INTERVIEW_FAILURE           = 5002


func GetInterviewFailed(w http.ResponseWriter, e error) {
    s := fmt.Sprint(MSG_GET_INTERVIEW_ERROR, " : ", e)
    w.WriteHeader(http.StatusInternalServerError);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: SERVERERROR_GET_INTERVIEW_FAILURE, Message: s })
}

func SaveInterviewFailed(w http.ResponseWriter, e error) {
    s := fmt.Sprint(MSG_SAVE_INTERVIEW_ERROR, " : ", e)
    w.WriteHeader(http.StatusInternalServerError);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: SERVERERROR_SAVE_INTERVIEW_FAILURE, Message: s })
}

func InterviewNotFound(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNotFound);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: NOTFOUND_INTERVIEW_NOTFOUND, Message: MSG_INTERVIEW_NOTFOUND })
}

func NoVersionProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_NOVERSION, MSG_NO_VERSION_PROVIDED)
}

func InvalidVersionProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_INVALIDVERSION, MSG_INVALID_VERSION)
}

func NoQueryParametersProvided(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_NOINPUTPARAMETERS, MSG_NO_PARAMETERS_PROVIDED)
}

func UnsupportedVersion(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_UNSUPPORTEDVERSION, MSG_UNSUPPORTED_VERSION)
}

func FailedDecodingBody(w http.ResponseWriter) {
    writeBadRequest(w, BADREQUEST_FAILED_DECODING_REQUEST_BODY, MSG_FAILED_READING_BODY)
}


func writeBadRequest(w http.ResponseWriter, code int, msg string) {
    w.WriteHeader(http.StatusBadRequest);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: code, Message: msg })
}
