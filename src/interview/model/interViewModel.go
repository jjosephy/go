package model

type InterviewModel struct {
    Candidate           string              `json:"candidate"`
    Comments            []CommentModel      `json:"comments"`
    Complete            bool                `json:"complete"`
    Id                  string              `json:"id"`
}
