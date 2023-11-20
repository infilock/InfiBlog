package main

import (
	"context"
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/infilock/InfiBlog/config"
	"github.com/infilock/InfiBlog/internal/repository/postgresql/pool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	location, err := time.LoadLocation(config.TwitterJobLocation.Get())
	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(location) //time.UTC

	//Every schedules a new periodic Job with an interval.
	_, errJobOne := scheduler.Every(1).Day().At(config.TwitterJobTime.Get()).Do(func() {
		log.Println("- start job Twitter ")

		dbCfg, err := config.GetDBConfig()
		if err != nil {
			log.Fatal("unable to get db config", err)

			return
		}

		database := config.ConnectionToPSQL(dbCfg)
		if err != nil {
			log.Fatal("unable to create database client", err)
		}

		ra := pool.NewArticleRepository(database)
		rsm := pool.NewSocialMediaRepository(database)

		articles, errArticles := ra.ListStatus(context.Background(), "publish")
		if errArticles != nil {
			log.Fatal(errArticles)
		}

		for i := 0; i < len(articles); i++ {
			tweet, errTweet := rsm.FindTweet(context.Background(), articles[i].QuestionID)
			if errTweet != nil {
				log.Fatal(errTweet)
			}

			//send on Twitter
			type Tweet struct {
				Text string `json:"text"`
			}

			t := Tweet{
				Text: tweet.Content,
			}

			c, errMarshal := json.Marshal(t)
			if errMarshal != nil {
				log.Fatal(errMarshal)
			}

			payload := strings.NewReader(string(c))

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			req, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, config.TwitterEndpointCreate.Get(), payload)
			if errRequest != nil {
				log.Fatal(errRequest)
			}

			req.Header.Add("Authorization",
				"OAuth "+
					"oauth_consumer_key=\""+config.TwitterOauthConsumerKey.Get()+"\","+
					"oauth_token=\""+config.TwitterOauthToken.Get()+"\","+
					"oauth_signature_method=\"HMAC-SHA1\","+
					"oauth_timestamp=\"1684073025\","+
					"oauth_nonce=\"Zzs5FDoFgtg\","+
					"oauth_version=\"1.0\","+
					"oauth_signature=\"ajkOdpVpe8VFBkvJ5MfWyR%2FKBfc%3D\"")
			req.Header.Add("Content-Type", "application/json")

			client := &http.Client{}
			res, errClient := client.Do(req)
			if errClient != nil {
				log.Fatal(errClient)
			}
			defer res.Body.Close()

			if strings.Contains(res.Status, "201") {
				//update status field on table social media
				if errUpdate := rsm.UpdateSocialStatus(context.Background(), tweet.ID); errUpdate != nil {
					log.Fatal(errUpdate)
				}
			}
		}
	})

	if errJobOne != nil {
		log.Fatal(errors.Wrap(errJobOne, "error on job Twitter"))
	}

	//StartBlocking starts all jobs and blocks the current thread.
	scheduler.StartBlocking()
}
