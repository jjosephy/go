package contract

import (
    "testing"
)

func IsTrue(t *testing.T, c bool, msg string) {
    if c != true {
        t.Fatal(msg)
    }
}

func Test_Success_CreateCommentContractV1(t *testing.T) {
    c := CommentContractV1 {
        Content: "This is a comment",
        Interviewer: "Interviewer",
        InterviewerId: "ID",
    }

    IsTrue(t, c.Content == "This is a comment", "Comment Content does not match")
    IsTrue(t, c.Interviewer == "Interviewer", "Interviewer does not match")
}
