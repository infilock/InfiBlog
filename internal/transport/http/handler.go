package http

import (
	"crawler/internal/service/article"
	"crawler/internal/service/question"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type handler struct {
	router      *mux.Router
	articleSvc  article.Service
	questionSvc question.Service
}

func NewHandler(
	articleSvc article.Service,
	questionSvc question.Service,
) http.Handler {
	router := mux.NewRouter()

	h := &handler{
		router:      router,
		articleSvc:  articleSvc,
		questionSvc: questionSvc,
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
