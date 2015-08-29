package model

type InterviewModel struct {
    Candidate   string
    Comments    []CommentModel
    Complete    bool
    Id          string
}
