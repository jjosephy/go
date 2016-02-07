package repository

import (
    "errors"
    "fmt"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "github.com/jjosephy/go/interview/model"
    "time"
)

// Interview Repository Interface //
type InterviewRepository interface {
    GetInterview(id string, name string) (model.InterviewModel, error)
    SaveInterview(model model.InterviewModel) (model.InterviewModel, error)
}

// DBInterviewRespository Implementation //
type DBInterviewRepository struct {
    DBSession *mgo.Session
    Uri string
}

func (r *DBInterviewRepository) GetConnection()(err error) {

    if r.Uri == "" {
        return errors.New("address for database not configured")
    }

    if r.DBSession != nil {
        return nil
    }

    // TODO: there is a bug here when session is closed we never reopen it. Need Retry
    timeout := 5 * time.Second
    r.DBSession, err = mgo.DialWithTimeout(r.Uri, timeout)
    if err != nil {
        fmt.Println("err trying to dial: ", err)
        return err
    }

    r.DBSession.SetMode(mgo.Monotonic, true)

    // TODO: Unique index on _id
    /*
    r.DBSession.SetMode(mgo.Monotonic, true)
    index := mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
	err = r.DBSession.DB("interview").C("interviews").EnsureIndex(index)
    if err != nil {
        return err
    }
    */
    return  nil
}

func(r *DBInterviewRepository) SaveInterview(m model.InterviewModel) (model.InterviewModel, error) {
    if err := r.GetConnection(); err != nil {
        return m, err
    }

    m.Id = bson.NewObjectId()
    if err := r.DBSession.DB("interview").C("interviews").Insert(&m); err != nil {
        return m, err
    }

    defer r.DBSession.Close()
    return m, nil
}

func(r *DBInterviewRepository) GetInterview(id string, name string) (m model.InterviewModel, err error) {
    //var m model.InterviewModel
    //var err error

    defer func() {

        if err != nil {
            fmt.Println("error is not nil %v", err)
        }

        e := recover()

        if e == nil {
            fmt.Println("Empty error in defer")
        } else {
            fmt.Printf("pkg:  %v", e)
            fmt.Println("")
        }


        switch x := e.(type) {
    		case string:
    			err = errors.New(x)
    		case error:
    			err = x
    		default:
    			//err = errors.New("Unknown panic in defer::GetInterview")
                //err = nil
		}

        //fmt.Println("err: ",  fmt.Println(err.(*errors.Error).ErrorStack()))

        //if r.DBSession != nil {
        //    r.DBSession.Close()
        //}

        // TODO: log failure
        //return nil, nil
    }()

    if err = r.GetConnection(); err != nil {
        return m, err
    }

    if id == "" && name == "" {
        return m, errors.New("invalid search params provided")
    }

    if valid := bson.IsObjectIdHex(id); valid == false {
        return m, errors.New("HexId")
    }

    // TODO: find by candidate name
    m = model.InterviewModel{}
    bid := bson.ObjectIdHex(id)
    err = r.DBSession.DB("interview").C("interviews").FindId(bid).One(&m)

    if err != nil {
        return m, err
    }

    return m, nil
}
