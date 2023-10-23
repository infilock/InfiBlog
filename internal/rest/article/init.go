package article

import (
	"github.com/infilock/InfiBlog/internal/service/article"
)

// handler .
type handler struct {
	articleSvc article.Service
}

// NewHandler .
func NewHandler(articleSvc article.Service) Contract {
	return &handler{
		articleSvc: articleSvc,
	}
}
