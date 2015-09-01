package repository

import (
    "errors"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "interview/model"
)

// Interview Repository Interface //
type InterviewRepository interface {
    GetInterview(id string, name string) (model.InterviewModel, error)
    SaveInterview(model model.InterviewModel) (error)
}

// DBInterviewRespository Implementation //
type DBInterviewRepository struct {
    DBSession *mgo.Session
    Uri string
}

func (r *DBInterviewRepository) CheckConnection()(error) {
    var err error

    if r.Uri == "" {
        return errors.New("address for database not configured")
    }
    if r.DBSession != nil {
        /*
        err = d.DBSession.Ping()
        if err == nil {
            return
        }
        */
        return nil
    }

    r.DBSession, err = mgo.Dial(r.Uri)
    if err != nil {
        return err
    }

    r.DBSession.SetMode(mgo.Monotonic, true)
    index := mgo.Index{
		Key:        []string{"queryid"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
	err = r.DBSession.DB("interview").C("interviews").EnsureIndex(index)

    if err != nil {
        defer r.DBSession.Close()
        return err
    }

    return  nil
}

func(r *DBInterviewRepository) SaveInterview(m model.InterviewModel) (error) {
    if err := r.CheckConnection(); err != nil {
        return err
    }

    return nil
}

func(r *DBInterviewRepository) GetInterview(id string, name string) (model.InterviewModel, error) {
    var m model.InterviewModel

    if err := r.CheckConnection(); err != nil {
        return m, err
    }

    if id == "" && name == "" {
        return m, errors.New("invalid search params provided")
    }

    m = model.InterviewModel{}
    err := r.DBSession.DB("interview").C("interviews").Find(bson.M{"queryid": id}).One(&m)

    if err != nil {
        return m, err
    }

    return m, nil
}
