package article

import (
	"context"
	dataModel "crawler/internal/repository/postgresql/article"
)

type Service interface {
	ListArticles(ctx context.Context) ([]dataModel.Entity, error)
	ListArticlesByStatus(ctx context.Context, status string) ([]dataModel.Entity, error)
}
