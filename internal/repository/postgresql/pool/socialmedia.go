package pool

import (
	"context"
	"database/sql"
	"github.com/infilock/InfiBlog/internal/service/socialmedia"
	"github.com/pkg/errors"
	"time"
)

const (
	queryListSocialByStatus = `SELECT id,question_id,type,content,status,created_at FROM socialmedias WHERE status=$1`
	queryUpdateSocialStatus = `UPDATE socialmedias SET status=1 WHERE id=$1`
	queryInsertSocialMedia  = `INSERT INTO socialmedias (question_id,type,content,created_at) VALUES ($1,$2,$3,$4)`
	queryFindTweet          = `SELECT id,question_id,type,content,status,created_at FROM socialmedias WHERE question_id=$1 AND type='twitter'`
)

type socialMediaRepo struct {
	db *sql.DB
}

func NewSocialMediaRepository(db *sql.DB) socialmedia.Repository {
	return &socialMediaRepo{db: db}
}

func (repo *socialMediaRepo) FindTweet(ctx context.Context, questionID string) (*socialmedia.Entity, error) {
	var m socialmedia.Entity

	r := repo.db.QueryRowContext(ctx, queryFindTweet, questionID)
	err := r.Scan(&m.ID, &m.QuestionID, &m.Type, &m.Content, &m.Status, &m.CreatedAt)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil //nolint:nilnil
	case err == nil:
		return &m, nil
	default:
		return nil, errors.Wrap(err, "error on scan row")
	}
}

func (repo *socialMediaRepo) UpdateSocialStatus(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, queryUpdateSocialStatus, id)
	if err != nil {
		return errors.Wrap(err, "error exec context db")
	}

	return nil
}

func (repo *socialMediaRepo) ListSocialByStatus(ctx context.Context, social, status string) ([]*socialmedia.Entity, error) {
	rows, err := repo.db.QueryContext(ctx, queryListSocialByStatus, social, status)

	defer func() { _ = rows.Close() }()

	items := make([]*socialmedia.Entity, 0)

	for rows.Next() {
		var m socialmedia.Entity

		if errScan := rows.Scan(&m.ID, &m.QuestionID, &m.Type, &m.Content, &m.Status, &m.CreatedAt); errScan != nil {
			return nil, errors.Wrap(errScan, "error on scan row")
		}

		items = append(items, &m)
	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "error on ListSocialByStatus: ")
	}

	return items, nil
}

func (repo *socialMediaRepo) Create(ctx context.Context, m []*socialmedia.Entity) error {
	for i := 0; i < len(m); i++ {
		args := []interface{}{
			m[i].QuestionID,
			m[i].Type,
			m[i].Content,
			time.Now(),
		}

		_, err := repo.db.ExecContext(ctx, queryInsertSocialMedia, args...)
		if err != nil {
			return errors.Wrap(err, "error on exec context for query insert")
		}
	}

	return nil
}
