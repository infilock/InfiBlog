package question

import (
	"github.com/infilock/InfiBlog/internal/service/question"
)

// handler .
type handler struct {
	questionSvc question.Service
}

// NewHandler .
func NewHandler(questionSvc question.Service) Contract {
	return &handler{
		questionSvc: questionSvc,
	}
}
