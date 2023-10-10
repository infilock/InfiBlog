package main

import (
	"context"
	cfg "crawler/internal/pkg/postgresql"
	repoArticle "crawler/internal/repository/postgresql/article"
	repoQuestion "crawler/internal/repository/postgresql/question"
	repoSocialMedia "crawler/internal/repository/postgresql/socialmedia"
	"github.com/go-co-op/gocron"
	"github.com/gosimple/slug"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/nasermirzaei89/env"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
	"time"
)

func main() {
	location, err := time.LoadLocation(env.MustGetString("LOCATION"))
	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(location) //time.UTC

	//Every schedules a new periodic Job with an interval.
	_, errJobOne := scheduler.Every(1).Day().At(env.MustGetString("JOB_TIME_QUESTION")).Do(func() {
		log.Println("- start job ChatGPT ")

		postgresql := cfg.GetPostgresqlConfig()
		defer func() { _ = postgresql.Close() }()

		rq := repoQuestion.NewQuestionRepository(postgresql)
		ra := repoArticle.NewArticleRepository(postgresql)
		rsm := repoSocialMedia.NewSocialMediaRepository(postgresql)

		res, errListStatus := rq.ListStatus(context.Background(), "0")
		if errListStatus != nil {
			log.Fatal(errListStatus)
		}

		for i := 0; i < len(res); i++ {
			client := openai.NewClient(env.MustGetString("GPT_API_KEY"))
			//API call to Create a completion for the chat message.
			resp, errChatCompletion := client.CreateChatCompletion(
				context.Background(),
				//ChatCompletionRequest represents a request structure for chat completion API.
				openai.ChatCompletionRequest{
					Model:       openai.GPT3Dot5Turbo, //GPT3 Defines the models provided by OpenAI to use when generating completions from OpenAI.
					Temperature: 0.2,
					//This composite literal allocates a new struct instance with the given values.
					Messages: []openai.ChatCompletionMessage{
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
			article := repoArticle.Entity{
				QuestionID: res[i].ID,
				Title:      title,
				Content:    content,
			}

			errCreate := ra.Create(context.Background(), article)
			if errCreate != nil {
				log.Fatal(errCreate, "error on create article user repo: ")
			}

			errUpdateArticle := rq.UpdateQuestionStatus(context.Background(), res[i].ID)
			if errUpdateArticle != nil {
				log.Println(errUpdateArticle)
				continue //nolint:nlreturn
			}
			//end: article

			//begin: social media
			socialmedia := []repoSocialMedia.Entity{
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

			errCreateSocialMedia := rsm.Create(context.Background(), socialmedia)
			if errCreateSocialMedia != nil {
				log.Fatal(errCreateSocialMedia, "error on create social media repo:")
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
	return content + "\n\n" + env.MustGetString("URL_BLOG_POST") + slug.Make(title)
}
