package http

import (
	"net/http"
)

func (h *handler) registerRoutes() {
	h.router.Methods(http.MethodGet).Path("/articles").HandlerFunc(h.HandlerListArticles())

	h.router.Methods(http.MethodPost).Path("/question").HandlerFunc(h.HandlerUploadQuestion())
	h.router.Methods(http.MethodGet).Path("/questions").HandlerFunc(h.HandlerListQuestions())
}
