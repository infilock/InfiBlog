package article

import (
	"context"
	repo "crawler/internal/repository/postgresql/article"
	"github.com/pkg/errors"
)

type service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (svc *service) ListArticles(ctx context.Context) ([]repo.Entity, error) {
	res, err := svc.repo.Lists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error on list articles service:")
	}

	return res, nil
}

func (svc *service) ListArticlesByStatus(ctx context.Context, status string) ([]repo.Entity, error) {
	res, err := svc.repo.ListStatus(ctx, status)
	if err != nil {
		return nil, errors.Wrap(err, "error on list articles by status service:")
	}

	return res, nil
}
