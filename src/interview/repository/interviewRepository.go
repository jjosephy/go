package repository

import (
    "errors"
    "gopkg.in/mgo.v2"
    "interview/model"
)

type InterviewRepository interface {
    GetInterview(id string, name string) (model.InterviewModel, error)
    SaveInterview(model model.InterviewModel) (error)
}

type DBInterviewRepository struct {
    DBSession *mgo.Session
    Uri string
}

func (r *DBInterviewRepository) checkConnection()(error) {
    var err error
    r.DBSession, err = mgo.Dial(r.Uri)
    if err != nil {
        return err
    }
    return  nil
}

func(r *DBInterviewRepository) SaveInterview(m model.InterviewModel) (error) {
    if err := r.checkConnection(); err != nil {
        return err
    }
    return nil
}

func(r *DBInterviewRepository) GetInterview(id string, name string) (model.InterviewModel, error) {
    var m model.InterviewModel

    if id == "" && name == "" {
        return m, errors.New("invalid search params provided")
    }

    comments := model.Comments {
        model.CommentModel { Content: "db Content", Interviewer: "interviewer 0", InterviewerId: "0" },
        model.CommentModel { Content: "db Content", Interviewer: "interviewer 1", InterviewerId: "1" },
        model.CommentModel { Content: "db Content", Interviewer: "interviewer 2", InterviewerId: "2" },
    }

    // Get a model and translate that
    m = model.InterviewModel {
        Candidate: "Candidate Name",
        Id: "hardcodedid",
        Comments: comments,
    }

    return m, nil
}
