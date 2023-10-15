package article

import "net/http"

type Contract interface {
	HandlerListArticles() http.HandlerFunc
}
