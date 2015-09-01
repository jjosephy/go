package model

import (
    "encoding/json"
    "testing"
)

func Test_Success_CreateErrorModel(t *testing.T) {
    m := ErrorModel {
        Message: "ErrMessage",
        ErrorCode: 1000,
    }

    if m.Message != "ErrMessage" {
        t.Fatal("failed")
    }
}

func Test_Success_SerializeCommentModel(t *testing.T) {
    // Get a model and translate that
    m := InterviewModel {
        Candidate: "Candidate Name",
        QueryId: 1,
        Comments: Comments {
            CommentModel { Content: "db Content", Interviewer: "interviewer 0", InterviewerId: "0" },
            CommentModel { Content: "db Content", Interviewer: "interviewer 1", InterviewerId: "1" },
            CommentModel { Content: "db Content", Interviewer: "interviewer 2", InterviewerId: "2" },
        },
    }

    _, err := json.Marshal(m)
    if err != nil {
        t.Fatalf("Failed Marshal model %v\n", err)
    }

    assertAreEqual(t, m.Candidate == "Candidate Name", "Candiate Names dont match")
}

func assertAreEqual(t *testing.T, c bool, msg string) {
    if c != true {
        t.Fatal(msg)
    }
}
