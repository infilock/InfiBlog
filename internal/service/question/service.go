package question

import (
	"context"
	"mime/multipart"
)

type Service interface {
	ListQuestions(ctx context.Context) ([]*Entity, error)
	CreateQuestion(ctx context.Context, file multipart.File, tagID, categoryID string) error
	ListQuestionsByStatus(ctx context.Context, status string) ([]*Entity, error)
}
