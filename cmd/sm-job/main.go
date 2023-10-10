package main

import (
	"context"
	cfg "crawler/internal/pkg/postgresql"
	repoArticle "crawler/internal/repository/postgresql/article"
	repoSocialMedia "crawler/internal/repository/postgresql/socialmedia"
	"encoding/json"
	"github.com/go-co-op/gocron"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/nasermirzaei89/env"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	location, err := time.LoadLocation(env.MustGetString("JOB_LOCATION_SM"))
	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(location) //time.UTC

	//Every schedules a new periodic Job with an interval.
	_, errJobOne := scheduler.Every(1).Day().At(env.MustGetString("JOB_TIME_SM")).Do(func() {
		log.Println("- start job Twitter ")

		postgresql := cfg.GetPostgresqlConfig()
		defer func() { _ = postgresql.Close() }()

		ra := repoArticle.NewArticleRepository(postgresql)
		rsm := repoSocialMedia.NewSocialMediaRepository(postgresql)

		articles, errArticles := ra.ListStatus(context.Background(), "publish")
		if errArticles != nil {
			log.Fatal(errArticles)
		}

		for i := 0; i < len(articles); i++ {
			tweet, errFindTweet := rsm.FindTweet(context.Background(), articles[i].QuestionID)
			if errFindTweet != nil {
				log.Fatal(errFindTweet)
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

			req, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, env.MustGetString("TW_ENDPOINT_CREATE"), payload)
			if errRequest != nil {
				log.Fatal(errRequest)
			}

			req.Header.Add("Authorization",
				"OAuth "+
					"oauth_consumer_key=\""+env.MustGetString("TW_OAUTH_CONSUMER_KEY")+"\","+
					"oauth_token=\""+env.MustGetString("TW_OAUTH_TOKEN")+"\","+
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
				errUpdateSocialStatus := rsm.UpdateSocialStatus(context.Background(), tweet.ID)
				if errUpdateSocialStatus != nil {
					log.Fatal(errUpdateSocialStatus)
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
