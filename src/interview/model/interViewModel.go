package model

type InterviewModel struct {
    Candidate           string      `json:"candidate"`
    Comments            string      `json:"comments"`
    Complete            bool        `json:"complete"`
    Id                  string      `json:"id"`
    Interviewer         string      `json:"interviewer"`
}
