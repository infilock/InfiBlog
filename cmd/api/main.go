package main

import (
	"context"
	"errors"
	"github.com/infilock/InfiBlog/config"
	"github.com/infilock/InfiBlog/internal/api"
	"github.com/infilock/InfiBlog/internal/repository/postgresql/pool"
	"github.com/infilock/InfiBlog/internal/service/article"
	"github.com/infilock/InfiBlog/internal/service/question"

	articleHndlr "github.com/infilock/InfiBlog/internal/rest/article"
	questioneHndlr "github.com/infilock/InfiBlog/internal/rest/question"

	//_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq" // import postgres driver
	"log"
	"net/http"

	"os"
	"os/signal"
	"time"
)

func main() {
	dbCfg, err := config.GetDBConfig()
	if err != nil {
		log.Fatal("unable to get db config: ", err)

		return
	}

	database := config.ConnectionToPSQL(dbCfg)
	if err != nil {
		log.Fatal("unable to create database client: ", err)
	}

	// define repository
	articleRepo := pool.NewArticleRepository(database)
	questionRepo := pool.NewQuestionRepository(database)

	// define services
	articleSvc := article.NewService(articleRepo)
	questionSvc := question.NewService(questionRepo)

	// define handler
	artHnd := articleHndlr.NewHandler(articleSvc)
	queHnd := questioneHndlr.NewHandler(questionSvc)

	services := api.NewHandler(artHnd, queHnd)

	// A Server defines parameters for running an HTTP server. The zero value for Server is a valid configuration.
	collectionServer := &http.Server{
		Addr:         "4030",
		WriteTimeout: time.Second * 15, // riteTimeout is the maximum duration before timing out writes of the response.
		ReadTimeout:  time.Second * 15, // eadTimeout is the maximum duration for reading the entire request, including the body.
		IdleTimeout:  time.Second * 60, // dleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled.
		Handler:      services,
	}

	go func() {
		log.Printf("Start application on listening: 4030")

		if errServe := collectionServer.ListenAndServe(); errServe != nil {
			if !errors.Is(errServe, http.ErrServerClosed) {
				log.Fatal(errServe)
			}
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	errShutdown := collectionServer.Shutdown(ctx)
	if errShutdown != nil {
		log.Fatal("error on Shutdown: ", errShutdown)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down...")
	os.Exit(0)
}
