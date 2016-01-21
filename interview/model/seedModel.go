package model

import (
    "gopkg.in/mgo.v2/bson"
)

type Seed struct {
    Increment   string
    Val         int
    Id          bson.ObjectId `bson:"_id,omitempty"`
}
