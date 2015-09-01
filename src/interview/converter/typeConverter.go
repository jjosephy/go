package converter

import (
    "encoding/json"
    "interview/contract/v1"
    "interview/model"
    "io"
    "strconv"
)

func DecodeContractFromBodyV1(r io.ReadCloser) (contract.InterviewContractV1, error) {
    decoder := json.NewDecoder(r)
    var c contract.InterviewContractV1
    err := decoder.Decode(&c)

    if err != nil {
        return c, err
    }

    return c, nil
}
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

    // TODO: validate success
    d := strconv.Itoa(m.QueryId)
    return contract.InterviewContractV1 {
        Id: d,
        Candidate: m.Candidate,
        Comments: comments,
    }
}

func ConvertContractToModelV1 (c contract.InterviewContractV1) (m model.InterviewModel) {
    // TODO: validate input

    comments := make([]model.CommentModel, len(c.Comments))

    for i := 0; i < len(c.Comments); i++ {
        mc := model.CommentModel  {
                Content: c.Comments[i].Content,
                Interviewer: c.Comments[i].Interviewer,
                InterviewerId: c.Comments[i].InterviewerId }
        comments[i] = mc
    }

    // TODO: validate success
    a, _:= strconv.Atoi(c.Id)

    return model.InterviewModel {
        QueryId: a,
        Candidate: c.Candidate,
        Comments: comments,
    }
}
