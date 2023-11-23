package config

import "os"

type Env string

func (e Env) Get() string {
	return os.Getenv(string(e))
}

// postgresql
//const (
//	PsqlHost Env = "PSQL_HOST"
//	PsqlPort Env = "PSQL_PORT"
//	PsqlUser Env = "PSQL_USER"
//	PsqlPass Env = "PSQL_PASS"
//	PsqlDB   Env = "PSQL_DB"
//)

// chat gpt
const (
	GptAPIKey      Env = "GPT_API_KEY"
	GptJobTime     Env = "GPT_JOB_TIME"
	GptJobLocation Env = "GPT_JOB_LOCATION"
)

// twitter
const (
	TwitterEndpointCreate   Env = "TWITTER_ENDPOINT_CREATE"
	TwitterOauthConsumerKey Env = "TWITTER_OAUTH_CONSUMER_KEY"
	TwitterOauthToken       Env = "TWITTER_OAUTH_TOKEN"
	TwitterJobTime          Env = "TWITTER_JOB_TIME"
	TwitterJobLocation      Env = "TWITTER_JOB_LOCATION"
)

// WordPress
const (
	UrlBlogPost               Env = "URL_BLOG_POST"
	WordPressUsername         Env = "WORDPRESS_USERNAME"
	WordPressPassword         Env = "WORDPRESS_PASSWORD"
	WordPressAPICreateArticle Env = "WORDPRESS_API_CREATE_ARTICLE"
	//WordPressAPICreateMedia   Env = "WORDPRESS_API_CREATE_MEDIA"
	WordPressAPICategories Env = "WORDPRESS_API_LIST_CATEGORIES"
	WordPressAPITag        Env = "WORDPRESS_API_LIST_TAG"
	WordPressJobTime       Env = "WORDPRESS_JOB_TIME"
	WordPressJobLocation   Env = "WORDPRESS_JOB_LOCATION"
)

type Psql struct {
	Driver string
	Host   string
	Port   string
	User   string
	Pass   string
	Name   string
}
