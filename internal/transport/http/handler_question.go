package http

import (
	dataModel "crawler/internal/repository/postgresql/question"
	"crawler/pkg/problem"
	"crawler/pkg/respond"
	"net/http"
	"path/filepath"
)

func (h *handler) HandlerUploadQuestion() http.HandlerFunc {
	hh := func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			respond.Done(w, r, problem.BadRequest("form file", err))

			return
		}

		extension := filepath.Ext(fileHeader.Filename)
		if extension != ".xlsx" {
			respond.Done(w, r, problem.BadRequest("The file extension must .xlsx", err))

			return
		}

		tagID := r.URL.Query().Get("tag_id")
		categoryID := r.URL.Query().Get("category_id")

		errCreateQuestion := h.questionSvc.CreateQuestion(r.Context(), file, tagID, categoryID)
		if errCreateQuestion != nil {
			respond.Done(w, r, problem.InternalServerError(errCreateQuestion))

			return
		}

		respond.Done(w, r, nil)

		return
	}

	return hh
}

func (h *handler) HandlerListQuestions() http.HandlerFunc {
	type Response struct {
		Results []dataModel.Entity `json:"results"`
	}

	hh := func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		if status == "0" {
			res, err := h.questionSvc.ListQuestionsByStatus(r.Context(), "0")
			if err != nil {
				respond.Done(w, r, problem.InternalServerError(err))

				return
			}

			respond.Done(w, r, Response{Results: res})

			return
		}

		if status == "1" {
			res, err := h.questionSvc.ListQuestionsByStatus(r.Context(), "1")
			if err != nil {
				respond.Done(w, r, problem.InternalServerError(err))

				return
			}

			respond.Done(w, r, Response{Results: res})

			return
		}

		res, err := h.questionSvc.ListQuestions(r.Context())
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(err))

			return
		}

		respond.Done(w, r, Response{Results: res})

		return
	}

	return hh
}
