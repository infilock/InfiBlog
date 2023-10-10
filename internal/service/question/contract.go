package question

import (
	"context"
	dataModel "crawler/internal/repository/postgresql/question"
	repo "crawler/internal/repository/postgresql/question"
	"mime/multipart"
)

type Service interface {
	ListQuestions(ctx context.Context) ([]dataModel.Entity, error)
	CreateQuestion(ctx context.Context, file multipart.File, tagID, categoryID string) error
	ListQuestionsByStatus(ctx context.Context, status string) ([]repo.Entity, error)
}
