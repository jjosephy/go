package repository

import (
    "errors"
)

type InterviewRepository interface {
    GetInterview(id string, name string) string
}

type MockInterviewRepository struct {
}

func(r *MockInterviewRepository) GetInterview(id string, name string) string {
    return "Mock"
}


type DBInterviewRepository struct {
}

func(r *DBInterviewRepository) GetInterview(id string, name string) string {
    return "Mock"
}

const (
    MOCK = 0
    DB = 1
)

type RespositoryFactory struct {
    repo InterviewRepository
}

func GetRepository(i int)(InterviewRepository, error) {
    //var r InterviewRepository = nil
    switch i {
        case MOCK:
            return new(MockInterviewRepository), nil
        case DB:
            return new(DBInterviewRepository), nil
        default:
            // Need error here
            return nil, errors.New("Invalid Type")
    }
}
