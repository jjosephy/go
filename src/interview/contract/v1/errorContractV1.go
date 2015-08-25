package contract

type ErrorDetailV1 struct {
    Message    string       `json:"message"`
    Code       int          `json:"errorCode"`
}
