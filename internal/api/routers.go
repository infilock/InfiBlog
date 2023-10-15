package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// listEndPoints list load endpoint api.
func listEndPoints(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	path, _ := route.GetPathTemplate()
	methods, _ := route.GetMethods()

	fmt.Println(methods, "--->", path)

	return nil
}

func (h *handler) registerRoutes() {

	h.router.Methods(http.MethodGet).Path("/articles").HandlerFunc(h.articleCtr.HandlerListArticles())

	h.router.Methods(http.MethodPost).Path("/question").HandlerFunc(h.questionCtr.HandlerUploadQuestion())
	h.router.Methods(http.MethodGet).Path("/questions").HandlerFunc(h.questionCtr.HandlerListQuestions())

	if err := h.router.Walk(listEndPoints); err != nil {
		log.Fatal(err)
	}
}
