package article

import (
	"context"
	"github.com/pkg/errors"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (svc *service) ListArticles(ctx context.Context) ([]*Entity, error) {
	res, err := svc.repo.Lists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error on list articles service:")
	}

	return res, nil
}

func (svc *service) ListArticlesByStatus(ctx context.Context, status string) ([]*Entity, error) {
	res, err := svc.repo.ListStatus(ctx, status)
	if err != nil {
		return nil, errors.Wrap(err, "error on list articles by status service:")
	}

	return res, nil
}
