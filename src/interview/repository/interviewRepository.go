package repository

import (
    "errors"
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "interview/model"
)

// Simple Data Access Layer //
type IDAL interface {
    GetInterview(id string, name string) (model.InterviewModel, error)
    SaveInterview(m model.InterviewModel) (error)
    CheckConnection()(error)
}

type DAL struct {
    DBSession *mgo.Session
    Uri string
}

func (d *DAL)GetInterview(id string, name string) (model.InterviewModel, error) {
    var m model.InterviewModel

    if err := d.CheckConnection(); err != nil {
        return m, err
    }

    if id == "" && name == "" {
        return m, errors.New("invalid search params provided")
    }

    m = model.InterviewModel{}
    err := d.DBSession.DB("interview").C("interviews").Find(bson.M{"Id": id}).One(&m)
	if err != nil {
		return m, err
	}

    fmt.Printf("in %v \n", m)

    return m, nil
}

func (d *DAL)SaveInterview(m model.InterviewModel) (error) {
    if err := d.CheckConnection(); err != nil {
        return err
    }

    return nil
}

func (d *DAL) CheckConnection()(error) {
    var err error
    if d.DBSession != nil {
        /*
        err = d.DBSession.Ping()
        if err == nil {
            return
        }
        */
        return nil
    }

    d.DBSession, err = mgo.Dial(d.Uri)
    if err != nil {
        return err
    }

    return  nil
}
// End Data Access Layer //

// Interview Repository Interface //
type InterviewRepository interface {
    GetInterview(id string, name string) (model.InterviewModel, error)
    SaveInterview(model model.InterviewModel) (error)
}

// DBInterviewRespository Implementation //
type DBInterviewRepository struct {
    Dal *DAL
}

func CreateDBInterviewRespository(dal *DAL)(DBInterviewRepository, error) {
    var d DBInterviewRepository
    if dal == nil {
        return d, errors.New("Invalid Arg: dal")
    }

    d.Dal = dal
    return d, nil
}

func(r *DBInterviewRepository) SaveInterview(m model.InterviewModel) (error) {
    return r.Dal.SaveInterview(m)
}

func(r *DBInterviewRepository) GetInterview(id string, name string) (model.InterviewModel, error) {
    m, err := r.Dal.GetInterview(id, name)

    if err != nil {
        return m, err
    }

    return m, nil
}
