package question

import "github.com/infilock/InfiBlog/internal/service/question"

type Response struct {
	Results []*question.Entity `json:"results"`
}
