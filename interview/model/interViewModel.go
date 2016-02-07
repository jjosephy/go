package model

import (
    "labix.org/v2/mgo/bson"
)

type InterviewModel struct {
    Candidate           string
    Comments            []CommentModel
    Complete            bool
    Id                  bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
