package contract

type CommentContractV1  struct {
    Content             string      `json:"content"`
    Interviewer         string      `json:"interviewer"`
    InterviewerId       string      `json:"interviewerId"`
}

type CommentsV1 []CommentContractV1
