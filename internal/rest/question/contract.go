package question

import "net/http"

type Contract interface {
	HandlerUploadQuestion() http.HandlerFunc
	HandlerListQuestions() http.HandlerFunc
}
