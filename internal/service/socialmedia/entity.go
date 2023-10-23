package socialmedia

import "time"

// Entity data base model for user entity.
type Entity struct {
	ID         string    `json:"id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	QuestionID string    `json:"question_id"`
	Type       string    `json:"type"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
}
