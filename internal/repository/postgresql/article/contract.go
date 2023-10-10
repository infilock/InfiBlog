package article

import (
	"context"
	"time"
)

type Repository interface {
	Lists(ctx context.Context) ([]Entity, error)
	Create(ctx context.Context, entity Entity) error
	ListStatus(ctx context.Context, status string) ([]Entity, error)
	UpdateArticleStatus(ctx context.Context, id string) error
}

// Entity data base model for user entity.
type Entity struct {
	ID         string    `json:"id,omitempty"`
	QuestionID string    `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
}
