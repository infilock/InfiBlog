# crawler with ChatGPT

## ChatGPT
ChatGPT is an artificial intelligence chatbot developed by OpenAI and released in November 2022. It is built on top of OpenAI's GPT-3.5 and GPT-4 foundational large language models and has been fine-tuned using both supervised and reinforcement learning techniques

## crawler
A Web crawler, sometimes called a spider or spiderbot and often shortened to crawler, is an Internet bot that systematically browses the World Wide Web and that is typically operated by search engines for the purpose of Web indexing

## **Installation requirements**

### SQL schema migration tool for Go
[github sql-migrate ](https://github.com/rubenv/sql-migrate)

To install the library and command line program, use the following:\
```go get -v github.com/rubenv/sql-migrate/...```

For Go version from 1.18, use:\
```go install github.com/rubenv/sql-migrate/...@latest```

### Go program configuration
```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
export GOBIN=$GOPATH/bin
```
## set api key chat gpt in file env
In the paths: ```cmd/api/```<br>
Create a https://platform.openai.com/ account for test and receive keys: ```GPT_API_KEY``` and ```GPT_ORG_ID``` set file env

## set api Twitter in file env
In the paths: ```cmd/sm-job/``` </br>
Create a Twitter account for test and receive keys: ```OAUTH_CONSUMER_KEY``` and ```OAUTH_TOKEN``` set file env

## set api Google Programmable Search Engine in file env
Google Programmable Search Engine is a platform provided by Google that allows web developers to feature specialized information in web searches, refine and categorize queries and create customized search engines, based on Google Search.<br>

* In the paths: ```cmd/wp-job/``` </br>
* Create a Twitter account for test and receive keys: ```GOOGLE_API_KEY``` and ```GOOGLE_SEARCH_ID``` set file env

## set user&pass blog WordPress in file env
WordPress is a web content management system. It was originally created as a tool to publish blogs but has evolved to support publishing other web content, including more traditional websites, mailing lists and Internet forum, media galleries, membership sites, learning management systems and online stores.<br>

* In the paths: ```cmd/wp-job/``` </br>
* set username and password account WordPress: ```WP_USERNAME``` and ```WP_PASSWORD``` set file env


## How to launch the program
**note:** To start working with API and jobs, you must change the name of the `.env.example` file to `.env`
### Step 1: Start Docker services
```shell
cd scripts
sh deploy.sh -ud
```
**note:** Down Docker services
```shell
cd scripts
sh deploy.sh -dd
```

### Creating a binary file for each service
#### for api `make api`
#### for gpt `make build-job-gpt`
#### for wp `make build-job-wp`
#### for sm `make build-job-sm`
**note:** All the created binaries are in path ` /bin` and you need to transfer them to path `/usr/local/bin/infiBlog`

### Step 2: running and config APIs
```shell
cd cmd/api
```
#### * Copy the `.env` file to `.env.api`
#### * Copy the service file from path `/scripts/services/infiBlog-api.service` to path `/etc/systemd/system`
#### * move the env.api file to `/usr/local/bin/infiBlog`
#### * move the binary file api from path /bin/api to path `/usr/local/bin/infiBlog`
## Run jobs

### job chat gpt
```shell
cd cmd/job-gpt
```
#### * Copy the `.env` file to `.env.gpt`
#### * Copy the service file from path `/scripts/services/infiBlog-gpt.service` to path `/etc/systemd/system`
#### * move the env.gpt file to `/usr/local/bin/infiBlog`
#### * move the binary file gpt_job from path /bin/gpt_job to path `/usr/local/bin/infiBlog`

### job wordpress blog
```shell
cd cmd/job-wp
```
#### * Copy the `.env` file to `.env.wp`
#### * Copy the service file from path `/scripts/services/infiBlog-wp.service` to path `/etc/systemd/system`
#### * move the env.wp file to `/usr/local/bin/infiBlog`
#### * move the binary file wp_job from path /bin/wp_job to path `/usr/local/bin/infiBlog`

### job social media
```shell
cd cmd/job-social-media
```
#### * Copy the `.env` file to `.env.sm`
#### * Copy the service file from path `/scripts/services/infiBlog-sm.service` to path `/etc/systemd/system`
#### * move the env.sm file to `/usr/local/bin/infiBlog`
#### * move the binary file sm_job from path /bin/sm_job to path `/usr/local/bin/infiBlog`

## Uploading questions and how to work
#### Open Postman and send the Excel file of the questions in the API upload question
#### Example of IP upload question or curl
```shell
curl --location '127.0.0.1:4030/question?tag_id=1,16&category_id=5,6' \
--form 'file=@"docs/help/questios.xlsx"'
```
#### Path of sample Excel file of questions
``docs/help/questios.xlsx``