package repository

import (
)

type InterviewRepository interface {
    GetInterview(id string, name string) string
}

type DBInterviewRepository struct {
}

func(r *DBInterviewRepository) GetInterview(id string, name string) string {
    return "Mock"
}
