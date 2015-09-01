package model

import (
    "gopkg.in/mgo.v2/bson"
)

type InterviewModel struct {
    Candidate           string
    Comments            []CommentModel
    Complete            bool
    QueryId             int //TODO: make this a uuid
    Id                  bson.ObjectId `bson:"_id,omitempty"`
}
