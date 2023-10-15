package article

import (
	"github.com/infilock/InfiBlog/pkg/res"
	"net/http"
)

func (h *handler) HandlerListArticles() http.HandlerFunc {
	hh := func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		if status == "" {
			rr, err := h.articleSvc.ListArticles(r.Context())
			if err != nil {
				res.Done(w, r, res.InternalServerError(err))

				return
			}

			res.Done(w, r, Response{Results: rr})

			return
		}

		if status == "draft" {
			rr, err := h.articleSvc.ListArticlesByStatus(r.Context(), status)
			if err != nil {
				res.Done(w, r, res.InternalServerError(err))

				return
			}

			res.Done(w, r, Response{Results: rr})

			return
		}

		if status == "publish" {
			rr, err := h.articleSvc.ListArticlesByStatus(r.Context(), status)
			if err != nil {
				res.Done(w, r, res.InternalServerError(err))

				return
			}

			res.Done(w, r, Response{Results: rr})

			return
		}
	}

	return hh
}
