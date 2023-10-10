package article

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

const (
	queryListArticle         = `SELECT id,question_id,title,content,status,created_at FROM articles`
	queryListArticleStatus   = `SELECT id,question_id,title,content,status,created_at FROM articles WHERE status=$1`
	queryUpdateArticleStatus = `UPDATE articles SET status='publish' WHERE id=$1`
	queryInsertArticle       = `INSERT INTO articles (question_id,title,content,created_at) VALUES ($1,$2,$3,$4)`
)

type articleRepo struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) Repository {
	return &articleRepo{db: db}
}

func (repo *articleRepo) UpdateArticleStatus(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, queryUpdateArticleStatus, id)
	if err != nil {
		return errors.Wrap(err, "error exec context db")
	}

	return nil
}

func (repo *articleRepo) ListStatus(ctx context.Context, status string) ([]Entity, error) {
	rows, err := repo.db.QueryContext(ctx, queryListArticleStatus, status)

	defer func() { _ = rows.Close() }()

	items := make([]Entity, 0)

	for rows.Next() {
		var m Entity

		//Type NullString is a new type with the underlying type string. These types are different.
		var _ NullString

		if errScan := rows.Scan(&m.ID, &m.QuestionID, &m.Title, &m.Content, &m.Status, &m.CreatedAt); errScan != nil {
			return nil, errors.Wrap(errScan, "error on scan row")
		}

		items = append(items, m)
	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "error on queryListArticle: ")
	}

	return items, nil
}

func (repo *articleRepo) Lists(ctx context.Context) ([]Entity, error) {
	rows, err := repo.db.QueryContext(ctx, queryListArticle)

	defer func() { _ = rows.Close() }()

	items := make([]Entity, 0)

	for rows.Next() {
		var m Entity

		//Type NullString is a new type with the underlying type string. These types are different.
		var _ NullString

		if errScan := rows.Scan(&m.ID, &m.QuestionID, &m.Title, &m.Content, &m.Status, &m.CreatedAt); errScan != nil {
			return nil, errors.Wrap(errScan, "error on scan row")
		}

		items = append(items, m)
	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "error on queryListArticle: ")
	}

	return items, nil
}

func (repo *articleRepo) Create(ctx context.Context, m Entity) error {
	args := []interface{}{
		m.QuestionID,
		m.Title,
		m.Content,
		time.Now(),
	}
	//name,last_name,email,password,user_name,category,gender,created_at
	_, err := repo.db.ExecContext(ctx, queryInsertArticle, args...)
	if err != nil {
		return errors.Wrap(err, "error on exec context for query insert")
	}

	return nil
}
