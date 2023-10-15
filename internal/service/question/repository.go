package question

import "context"

type Repository interface {
	Lists(ctx context.Context) ([]*Entity, error)
	Create(ctx context.Context, entity Entity) error
	ListStatus(ctx context.Context, status string) ([]*Entity, error)
	UpdateQuestionStatus(ctx context.Context, id string) error
	Find(ctx context.Context, id string) (*Entity, error)
}
