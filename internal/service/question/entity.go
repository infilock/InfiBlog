package question

import "time"

// Entity data base model for user entity.
type Entity struct {
	ID         string    `json:"id"`
	Question   string    `json:"question"`
	Rule       string    `json:"rule"`
	CategoryID string    `json:"category_id"`
	TagID      string    `json:"tag_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
