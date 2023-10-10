package question

import (
	"context"
	"time"
)

type Repository interface {
	Lists(ctx context.Context) ([]Entity, error)
	Create(ctx context.Context, entity Entity) error
	ListStatus(ctx context.Context, status string) ([]Entity, error)
	UpdateQuestionStatus(ctx context.Context, id string) error
	Find(ctx context.Context, id string) (*Entity, error)
}

// Entity data base model for user entity.
type Entity struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Question   string    `json:"question"`
	Rule       string    `json:"rule"`
	CategoryID string    `json:"category_id"`
	TagID      string    `json:"tag_id"`
	Status     string    `json:"status"`
}
