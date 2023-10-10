package http

import (
	"crawler/internal/repository/postgresql/article"
	"crawler/pkg/problem"
	"crawler/pkg/respond"
	"net/http"
)

func (h *handler) HandlerListArticles() http.HandlerFunc {
	type Response struct {
		Results []article.Entity `json:"results"`
	}

	hh := func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		if status == "" {
			res, err := h.articleSvc.ListArticles(r.Context())
			if err != nil {
				respond.Done(w, r, problem.InternalServerError(err))

				return
			}

			respond.Done(w, r, Response{Results: res})

			return
		}

		if status == "draft" {
			res, err := h.articleSvc.ListArticlesByStatus(r.Context(), status)
			if err != nil {
				respond.Done(w, r, problem.InternalServerError(err))

				return
			}

			respond.Done(w, r, Response{Results: res})

			return
		}

		if status == "publish" {
			res, err := h.articleSvc.ListArticlesByStatus(r.Context(), status)
			if err != nil {
				respond.Done(w, r, problem.InternalServerError(err))

				return
			}

			respond.Done(w, r, Response{Results: res})

			return
		}
	}

	return hh
}
