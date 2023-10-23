package article

import "github.com/infilock/InfiBlog/internal/service/article"

type Response struct {
	Results []*article.Entity `json:"results"`
}
