package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/infilock/InfiBlog/internal/rest/article"
	"github.com/infilock/InfiBlog/internal/rest/question"
	"log"
	"net/http"
	"os"
)

type handler struct {
	router      *mux.Router
	articleCtr  article.Contract
	questionCtr question.Contract
}

func NewHandler(
	articleCtr article.Contract,
	questionCtr question.Contract,
) http.Handler {
	router := mux.NewRouter()

	h := &handler{
		router:      router,
		articleCtr:  articleCtr,
		questionCtr: questionCtr,
	}

	// register routes
	h.registerRoutes()

	return h
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var hh http.Handler

	hh = h.router
	hh = h.Logger(log.New(os.Stdout, fmt.Sprintln(), 0))(hh)
	hh = h.RecoverPanic()(hh)

	hh.ServeHTTP(w, r)
}
