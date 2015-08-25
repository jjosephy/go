package contract

type InterviewContractV1    struct {
    Candidate               string                      `json:"candidate"`
    Comments                []CommentContractV1         `json:"comments"`
    Complete                bool                        `json:"complete"`
    Id                      string                      `json:"id"`
    Interviewer             string                      `json:"interviewer"`
}

type InterviewsV1 []InterviewContractV1
