package socialmedia

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, entity []Entity) error
	ListSocialByStatus(ctx context.Context, social /* linkedin or twitter */, status /* 0 or 1 */ string) ([]Entity, error)
	FindTweet(ctx context.Context, questionID string) (*Entity, error)
	UpdateSocialStatus(ctx context.Context, id string) error
}

// Entity data base model for user entity.
type Entity struct {
	ID         string    `json:"id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	QuestionID string    `json:"question_id"`
	Type       string    `json:"type"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
}
