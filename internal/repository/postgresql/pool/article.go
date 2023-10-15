package pool

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/infilock/InfiBlog/internal/service/article"
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

func NewArticleRepository(db *sql.DB) article.Repository {
	return &articleRepo{db: db}
}

func (repo *articleRepo) UpdateArticleStatus(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, queryUpdateArticleStatus, id)
	if err != nil {
		return errors.Wrap(err, "error exec context db")
	}

	return nil
}

func (repo *articleRepo) ListStatus(ctx context.Context, status string) ([]*article.Entity, error) {
	rows, err := repo.db.QueryContext(ctx, queryListArticleStatus, status)

	defer func() { _ = rows.Close() }()

	items := make([]*article.Entity, 0)

	for rows.Next() {
		var m article.Entity

		//Type NullString is a new type with the underlying type string. These types are different.
		var _ NullString

		if errScan := rows.Scan(&m.ID, &m.QuestionID, &m.Title, &m.Content, &m.Status, &m.CreatedAt); errScan != nil {
			return nil, errors.Wrap(errScan, "error on scan row")
		}

		items = append(items, &m)
	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "error on queryListArticle: ")
	}

	return items, nil
}

func (repo *articleRepo) Lists(ctx context.Context) ([]*article.Entity, error) {
	rows, err := repo.db.QueryContext(ctx, queryListArticle)

	defer func() { _ = rows.Close() }()

	items := make([]*article.Entity, 0)

	for rows.Next() {
		var m article.Entity

		//Type NullString is a new type with the underlying type string. These types are different.
		var _ NullString

		if errScan := rows.Scan(&m.ID, &m.QuestionID, &m.Title, &m.Content, &m.Status, &m.CreatedAt); errScan != nil {
			return nil, errors.Wrap(errScan, "error on scan row")
		}

		items = append(items, &m)
	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "error on queryListArticle: ")
	}

	return items, nil
}

func (repo *articleRepo) Create(ctx context.Context, m article.Entity) error {
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

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""

		return nil
	}

	strVal, ok := value.(string)
	if !ok {
		return errors.New("Column is not a string")
	}

	*s = NullString(strVal)

	return nil
}

func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}

	return string(s), nil
}
