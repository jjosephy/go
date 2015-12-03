package httperror

import (
    "encoding/json"
    "fmt"
    "interview/model"
    "net/http"
)

const (
    MSG_NO_VERSION_PROVIDED               = "No Version Provided"
    MSG_INVALID_VERSION                   = "Invalid Version Provided"
    MSG_NO_PARAMETERS_PROVIDED            = "No Parameters Provided"
    MSG_UNSUPPORTED_VERSION               = "Unsupported Version Provided"
    MSG_GET_INTERVIEW_ERROR               = "Error occurred tyring to Get interview details"
    MSG_SAVE_INTERVIEW_ERROR              = "Error occurred trying to Save the interview"
    MSG_FAILED_READING_BODY               = "Failed to parse the request body"
    MSG_INTERVIEW_NOTFOUND                = "Interview not found"
    MSG_INTERVIEW_INVALIDID               = "Invalid Interview Id"
    MSG_UNAUTHORIZED                      = "Unauthorized"
)

// 400 Errors
const BADREQUEST_NOINPUTPARAMETERS                  = 3000
const BADREQUEST_NOVERSION                          = 3001
const BADREQUEST_INVALIDVERSION                     = 3002
const BADREQUEST_UNSUPPORTEDVERSION                 = 3003
const BADREQUEST_FAILED_DECODING_REQUEST_BODY       = 3004
const BADREQUEST_INTERVIEW_INVALIDID                = 3005

// 404 Errors
const NOTFOUND_INTERVIEW_NOTFOUND                   = 4000

//401 Errors
const NOAUTHORIZED_LDAPCONNECT_FAILURE             = 6000
const NOAUTHORIZED_FAILURE                         = 6001

// 500 Errors
const SERVERERROR_GENERAL                          = 5000
const SERVERERROR_GET_INTERVIEW_FAILURE            = 5001
const SERVERERROR_SAVE_INTERVIEW_FAILURE           = 5002

func Unauthorized(w http.ResponseWriter) {
    w.WriteHeader(http.StatusUnauthorized);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: NOAUTHORIZED_FAILURE, Message: MSG_UNAUTHORIZED })
}

func AuthConnectToLDAPFailure(w http.ResponseWriter) {
    w.WriteHeader(http.StatusUnauthorized);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: NOAUTHORIZED_LDAPCONNECT_FAILURE, Message: MSG_UNAUTHORIZED })
}

func GeneralServerError(w http.ResponseWriter, msg string) {
    w.WriteHeader(http.StatusInternalServerError);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: SERVERERROR_GENERAL, Message: msg })
}

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

func InvalidInterviewId(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNotFound);
    json.NewEncoder(w).Encode(model.ErrorModel { ErrorCode: BADREQUEST_INTERVIEW_INVALIDID, Message: MSG_INTERVIEW_INVALIDID })
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
