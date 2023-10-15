package article

import "time"

// Entity data base model for user entity.
type Entity struct {
	ID         string    `json:"id,omitempty"`
	QuestionID string    `json:"question_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
