package model

type CommentModel struct {
    Content         string
    Interviewer     string
    InterviewerId   int
}

type Comments []CommentModel
