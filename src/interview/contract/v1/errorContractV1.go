package contract

type ErrorDetail struct {
    Message    string       `json:"message"`
    Code       int          `json:"errorCode"`
}
