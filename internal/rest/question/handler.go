package question

import (
	"fmt"
	"github.com/infilock/InfiBlog/pkg/res"
	"net/http"
	"path/filepath"
)

func (h *handler) HandlerUploadQuestion() http.HandlerFunc {
	hh := func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			res.Done(w, r, res.BadRequest(fmt.Sprint(err)))

			return
		}

		extension := filepath.Ext(fileHeader.Filename)
		if extension != ".xlsx" {
			res.Done(w, r, res.BadRequest("The file extension must .xlsx"))

			return
		}

		tagID := r.URL.Query().Get("tag_id")
		categoryID := r.URL.Query().Get("category_id")

		errCreateQuestion := h.questionSvc.CreateQuestion(r.Context(), file, tagID, categoryID)
		if errCreateQuestion != nil {
			res.Done(w, r, res.InternalServerError(errCreateQuestion))

			return
		}

		res.Done(w, r, nil)

		return
	}

	return hh
}

func (h *handler) HandlerListQuestions() http.HandlerFunc {
	hh := func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		if status == "0" {
			rr, err := h.questionSvc.ListQuestionsByStatus(r.Context(), "0")
			if err != nil {
				res.Done(w, r, res.InternalServerError(err))

				return
			}

			res.Done(w, r, Response{Results: rr})

			return
		}

		if status == "1" {
			rr, err := h.questionSvc.ListQuestionsByStatus(r.Context(), "1")
			if err != nil {
				res.Done(w, r, res.InternalServerError(err))

				return
			}

			res.Done(w, r, Response{Results: rr})

			return
		}

		rr, err := h.questionSvc.ListQuestions(r.Context())
		if err != nil {
			res.Done(w, r, res.InternalServerError(err))

			return
		}

		res.Done(w, r, Response{Results: rr})

		return
	}

	return hh
}
