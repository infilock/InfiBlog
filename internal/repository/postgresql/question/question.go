package question

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

const (
	queryListQuestion         = `SELECT question,status,category_id,tag_id,created_at FROM questions`
	queryListQuestionStatus   = `SELECT id,question,rule,category_id,tag_id,status,created_at FROM questions WHERE status=$1`
	queryUpdateQuestionStatus = `UPDATE questions SET status=1 WHERE id=$1`
	queryInsertQuestion       = `INSERT INTO questions (question,rule,category_id,tag_id,created_at) VALUES ($1,$2,$3,$4,$5)`
	queryFindQuestion         = `SELECT id,question,rule,category_id,tag_id,status,created_at FROM questions WHERE id=$1`
)

type questionRepo struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) Repository {
	return &questionRepo{db: db}
}

func (repo *questionRepo) UpdateQuestionStatus(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, queryUpdateQuestionStatus, id)
	if err != nil {
		return errors.Wrap(err, "error exec context db")
	}

	return nil
}

func (repo *questionRepo) ListStatus(ctx context.Context, status string) ([]Entity, error) {
	rows, err := repo.db.QueryContext(ctx, queryListQuestionStatus, status)

	defer func() { _ = rows.Close() }()

	items := make([]Entity, 0)

	for rows.Next() {
		var m Entity

		if errScan := rows.Scan(&m.ID, &m.Question, &m.Rule, &m.CategoryID, &m.TagID, &m.Status, &m.CreatedAt); errScan != nil {
			return nil, errors.Wrap(errScan, "error on scan row")
		}

		items = append(items, m)
	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "error on queryListQuestion: ")
	}

	return items, nil
}

func (repo *questionRepo) Lists(ctx context.Context) ([]Entity, error) {
	rows, err := repo.db.QueryContext(ctx, queryListQuestion)

	defer func() { _ = rows.Close() }()

	items := make([]Entity, 0)

	for rows.Next() {
		var m Entity

		if errScan := rows.Scan(&m.ID, &m.Question, &m.Rule, &m.CategoryID, &m.TagID, &m.Status, &m.CreatedAt); errScan != nil {
			return nil, errors.Wrap(errScan, "error on scan row")
		}

		items = append(items, m)
	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "error on queryListQuestion: ")
	}

	return items, nil
}

func (repo *questionRepo) Create(ctx context.Context, m Entity) error {
	args := []interface{}{
		m.Question,
		m.Rule,
		m.CategoryID,
		m.TagID,
		time.Now(),
	}

	_, err := repo.db.ExecContext(ctx, queryInsertQuestion, args...)
	if err != nil {
		return errors.Wrap(err, "error on exec context for query insert")
	}

	return nil
}

func (repo *questionRepo) Find(ctx context.Context, id string) (*Entity, error) {
	var m Entity

	r := repo.db.QueryRowContext(ctx, queryFindQuestion, id)
	err := r.Scan(&m.ID, &m.Question, &m.Rule, &m.CategoryID, &m.TagID, &m.Status, &m.CreatedAt)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil //nolint:nilnil
	case err == nil:
		return &m, nil
	default:
		return nil, errors.Wrap(err, "error on scan row")
	}
}
