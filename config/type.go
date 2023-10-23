package config

import "os"

type Env string

func (e Env) Get() string {
	return os.Getenv(string(e))
}

// postgresql
const (
	PsqlHost         Env = "PSQL_HOST"
	PsqlPort         Env = "PSQL_PORT"
	PsqlUser         Env = "PSQL_USER"
	PsqlPass         Env = "PSQL_PASS"
	PsqlDB           Env = "PSQL_DB"
	PsqlMigrationDir Env = "PSQL_MIGRATION_DIR"
)

// chat gpt
const (
	GptAPIKey      Env = "GPT_API_KEY"
	GptOrgID       Env = "GPT_ORG_ID"
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

// config app
const (
	PortApp    Env = "APP_PORT"
	UploadPath Env = "UPLOAD_PATH"
)

// WordPress
const (
	UrlBlogPost               Env = "URL_BLOG_POST"
	WordPressUsername         Env = "WORDPRESS_USERNAME"
	WordPressPassword         Env = "WORDPRESS_PASSWORD"
	WordPressAPICreateArticle Env = "WORDPRESS_API_CREATE_ARTICLE"
	WordPressAPICreateMedia   Env = "WORDPRESS_API_CREATE_MEDIA"
	WordPressJobTime          Env = "WORDPRESS_JOB_TIME"
	WordPressJobLocation      Env = "WORDPRESS_JOB_LOCATION"
)

// google search engine
const (
	GoogleAPIKEY   Env = "GOOGLE_API_KEY"
	GoogleSearchID Env = "GOOGLE_SEARCH_ID"
)

type Psql struct {
	Driver       string
	MigrationDir string
	Host         string
	Port         string
	User         string
	Pass         string
	Name         string
}
