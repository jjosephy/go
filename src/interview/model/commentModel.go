package model

type CommentModel struct {
    Content         string
    Interviewer     string
    InterviewerId   string
}

type Comments []CommentModel
