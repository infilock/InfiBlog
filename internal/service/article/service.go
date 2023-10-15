package article

import (
	"context"
)

type Service interface {
	ListArticles(ctx context.Context) ([]*Entity, error)
	ListArticlesByStatus(ctx context.Context, status string) ([]*Entity, error)
}
