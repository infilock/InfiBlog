package socialmedia

import "context"

type Repository interface {
	Create(ctx context.Context, entity []*Entity) error
	ListSocialByStatus(ctx context.Context, social /* linkedin or twitter */, status string) ([]*Entity, error) // status => 0 daft , 1 publish
	FindTweet(ctx context.Context, questionID string) (*Entity, error)
	UpdateSocialStatus(ctx context.Context, id string) error // change status from 0 to 1 send on social media
}
