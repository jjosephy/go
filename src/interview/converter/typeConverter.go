package converter

import (
    "interview/contract/v1"
    "interview/model"
)

func ConvertModelToContractV1 (m model.InterviewModel) (c contract.InterviewContractV1) {
    // TODO: validate input

    comments := make([]contract.CommentContractV1, len(m.Comments))

    for i := 0; i < len(m.Comments); i++ {
        cm := contract.CommentContractV1  {
                Content: m.Comments[i].Content,
                Interviewer: m.Comments[i].Interviewer,
                InterviewerId: m.Comments[i].InterviewerId }
        comments[i] = cm
    }

    return contract.InterviewContractV1 {
        Id: m.Id,
        Candidate: m.Candidate,
        Comments: comments,
    }
}
