package main

import (
	cfg "crawler/internal/pkg/postgresql"
	repoArticle "crawler/internal/repository/postgresql/article"
	repoQuestion "crawler/internal/repository/postgresql/question"
	"crawler/internal/service/article"
	"crawler/internal/service/question"
	router "crawler/internal/transport/http"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq" // import postgres driver
	"github.com/nasermirzaei89/env"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	postgresql := cfg.GetPostgresqlConfig()
	defer func() { _ = postgresql.Close() }()

	articleRepo := repoArticle.NewArticleRepository(postgresql)
	articleSvc := article.NewService(articleRepo)

	questionRepo := repoQuestion.NewQuestionRepository(postgresql)
	questionSvc := question.NewService(questionRepo)

	services := router.NewHandler(articleSvc, questionSvc)

	//A Signal represents an operating system signal.
	stop := make(chan os.Signal, 1)
	// Notify causes package signal to relay incoming signals to c.
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// A "go" statement starts the execution of a function call as an independent concurrent thread of control, or goroutine, within the same address space.
	go func() {
		log.Printf("Start application on listening on %s", env.MustGetString("APP_PORT"))

		// A Server defines parameters for running an HTTP server. The zero value for Server is a valid configuration.
		s := &http.Server{
			Addr:              env.MustGetString("APP_HOST") + env.MustGetString("APP_PORT"),
			ReadHeaderTimeout: 0,
			Handler:           services,
		}

		if err := s.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stop

	time.Sleep(1 * time.Second)
	log.Printf("shutting down application ...\n")
}
