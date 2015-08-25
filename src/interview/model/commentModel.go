package model

type CommentModel struct {
    Content         string      `json:"comments"`
    Interviewer     string      `json:"interviewer"`
    InterviewerId   int         `json:"interviewer"`
}

type Comments []CommentModel
