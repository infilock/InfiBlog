package main

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/gosimple/slug"
	"github.com/infilock/InfiBlog/config"
	"github.com/infilock/InfiBlog/internal/repository/postgresql/pool"
	"github.com/infilock/InfiBlog/internal/service/article"
	"github.com/infilock/InfiBlog/internal/service/socialmedia"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
	"time"
)

func main() {
	location, err := time.LoadLocation(config.GptJobLocation.Get())
	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(location) //time.UTC

	//Every schedules a new periodic Job with an interval.
	_, errJobOne := scheduler.Every(1).Day().At(config.GptJobTime.Get()).Do(func() {
		log.Println("- start job ChatGPT ")

		dbCfg, err := config.GetDBConfig()
		if err != nil {
			log.Fatal("unable to get db config", err)

			return
		}

		database := config.ConnectionToPSQL(dbCfg)
		if err != nil {
			log.Fatal("unable to create database client", err)
		}

		rq := pool.NewQuestionRepository(database)
		ra := pool.NewArticleRepository(database)
		rsm := pool.NewSocialMediaRepository(database)

		res, errStatus := rq.ListStatus(context.Background(), "pending")
		if errStatus != nil {
			log.Fatal(errStatus)
		}

		for i := 0; i < len(res); i++ {
			client := openai.NewClient(config.GptAPIKey.Get())
			//API call to Create a completion for the chat message.
			resp, errChatCompletion := client.CreateChatCompletion(context.Background(),
				//ChatCompletionRequest represents a request structure for chat completion API.
				openai.ChatCompletionRequest{
					Model:       openai.GPT3Dot5Turbo, //GPT3 Defines the models provided by OpenAI to use when generating completions from OpenAI.
					Temperature: 0.2,
					Messages: []openai.ChatCompletionMessage{ //This composite literal allocates a new struct instance with the given values.
						{
							// Chat message role defined by the OpenAI API
							Role:    openai.ChatMessageRoleUser,
							Content: res[i].Question + res[i].Rule,
						},
					},
				},
			)

			if errChatCompletion != nil {
				log.Fatal("error on call to create a completion for the chat message:", errChatCompletion)
			}

			title, content, linkedin, twitter := extractContent(resp.Choices[0].Message.Content)
			log.Println("title -- > : ", title)

			//begin: article
			itemArticle := article.Entity{
				QuestionID: res[i].ID,
				Title:      title,
				Content:    content,
			}

			if errArticle := ra.Create(context.Background(), itemArticle); errArticle != nil {
				log.Fatal(errArticle, "error on create article: ")
			}

			if errQuestion := rq.UpdateQuestionStatus(context.Background(), res[i].ID); errQuestion != nil {
				log.Println(errQuestion)

				continue
			}
			//end: article

			//begin: social media
			itemSocialMedia := []*socialmedia.Entity{
				{
					QuestionID: res[i].ID,
					Type:       "twitter",
					Content:    attachURL(title, twitter),
				},
				{
					QuestionID: res[i].ID,
					Type:       "linkedin",
					Content:    attachURL(title, linkedin),
				},
			}

			if errSocialMedia := rsm.Create(context.Background(), itemSocialMedia); errSocialMedia != nil {
				log.Fatal(errSocialMedia, "error on create social media repo:")
			}
			//end: social media
		}
	})

	if errJobOne != nil {
		log.Fatal(errors.Wrap(errJobOne, "error on job gpt"))
	}

	//StartBlocking starts all jobs and blocks the current thread.
	scheduler.StartBlocking()
}

func extractContent(value string) (string, string, string, string) {
	var title, content, linkedIn, twitter string

	posTitle := strings.Index(value, "Post Title: ")
	posLastTitle := strings.Index(value, "\n\n")
	posFirstAdjusted := posTitle + len("Post Title: ")
	title = value[posFirstAdjusted:posLastTitle]

	posContent := strings.Index(value[posLastTitle:], "\n\n")
	posLastContent := strings.Index(value[posLastTitle:], "\n\nLinkedIn Post: ")
	posFirstAdjusted1 := posContent + len("\n\n")
	content = value[posLastTitle:][posFirstAdjusted1:posLastContent]

	posLinkedIn := strings.Index(value[posLastContent:], "\n\nLinkedIn Post: ")
	posLastLinkedIn := strings.Index(value[posLastContent:], "\n\nTwitter Post: ")
	posFirstAdjusted10 := posLinkedIn + len("\n\nLinkedIn Post: ")
	linkedIn = value[posLastContent:][posFirstAdjusted10:posLastLinkedIn]

	posTwitter := strings.Index(value[posLastLinkedIn:], "\n\nTwitter Post: ")
	posFirstAdjustedTwitter := posTwitter + len("\n\nTwitter Post: ")
	twitter = value[posLastLinkedIn:][posFirstAdjustedTwitter:]

	return title, content, linkedIn, twitter
}

func attachURL(title, content string) string {
	return content + "\n\n" + config.UrlBlogPost.Get() + slug.Make(title)
}
