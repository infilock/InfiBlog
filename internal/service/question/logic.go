package question

import (
	"context"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (svc *service) ListQuestionsByStatus(ctx context.Context, status string) ([]*Entity, error) {
	res, err := svc.repo.ListStatus(ctx, status)
	if err != nil {
		return nil, errors.Wrap(err, "error on list question service:")
	}

	return res, nil
}

func (svc *service) ListQuestions(ctx context.Context) ([]*Entity, error) {
	res, err := svc.repo.Lists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error on list question service:")
	}

	return res, nil
}

func (svc *service) CreateQuestion(ctx context.Context, file multipart.File, tagID, categoryID string) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return errors.Wrap(err, "error on openReader read data stream:")
	}

	rows, errRows := f.GetRows("Sheet1")
	if errRows != nil {
		return errors.Wrap(errRows, "error on return all the rows:")
	}

	for j, row := range rows {
		if j == 0 {
			continue
		}

		m := Entity{
			Question:   row[0],
			Rule:       row[1],
			CategoryID: categoryID,
			TagID:      tagID,
		}

		errCreate := svc.repo.Create(ctx, m)
		if errCreate != nil {
			return errors.Wrap(errCreate, "error on create question service:")
		}
	}

	return nil
}
