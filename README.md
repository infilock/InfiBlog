# Infilock Blog generator
This GitHub project is a versatile tool designed to transform a set of keywords into professionally styled and standardized blog posts across various domains. It operates autonomously, effortlessly generating and publishing a specified number of blog posts on a regular schedule, requiring no human intervention. Moreover, the tool seamlessly publishes these posts in the correct format to WordPress, Instagram, LinkedIn, and other platforms.

## Setup Instruction

### 1- Install GO
Follow [Go installation steps](https://go.dev/doc/install).

### 2- Install SQL schema migration tool
Install the library and command line program\
```go get -v github.com/rubenv/sql-migrate/...```

Note: For Go version from 1.18, use:\
```go install github.com/rubenv/sql-migrate/...@latest```


### 3- Configure jobs
Configure ChatGPT, X(former twitter) and WordPress settings in .env.example file and rename the file to .env

- **ChatGPT Job:** This job involves querying chatGPT to create blog posts, which are then stored in a database.
Get  ```GPT_API_KEY``` and ```GPT_ORG_ID``` from [OpenAI](https://platform.openai.com/ ) account.

- **X platform Job:** This job involves scheduling the publication of blog posts to X (Twitter).
Get ```OAUTH_CONSUMER_KEY``` and ```OAUTH_TOKEN``` from an X account.

- **WordPress Job:** Similarly, this job publishes blogpost to a WordPress weblog.
Get ```WP_USERNAME``` and ```WP_PASSWORD``` from the WordPress.

**(Optional)** you can set ```GOOGLE_API_KEY``` and ```GOOGLE_SEARCH_ID``` in the environment file to autonomously search for an image for the weblog.


### 4- Run InfiBlog
- **Step 1:** Install PgAdmin and Postgres database
```shell
cd scripts
sh deploy.sh -ud
```

- **Step 2:** Build Binaries
```make infiBlog```

**note:** All the created binaries are stored in path `/bin`

- **Step 3:** Run Jobs
```
 Copy the service file from path `/scripts/services/infiBlog.service` to path `/etc/systemd/system`
 Move the .env file to `/usr/local/bin/infiBlog`
 Move the binary file from path `/bin/` to path `/usr/local/bin/infiBlog`
```


## Test
To simplify API calls, there are a Postman and a Swager file in the `docs` folder. There is also a sample Excel sheet in the `docs/help` path.
- Insert your required questions into the Excel file.
- Open Postman and upload the Excel file using the question API.
- (alternative) you can use curl command to upload questions.
```shell
curl --location '127.0.0.1:4030/question?tag_id=1,16&category_id=5,6' \
--form 'file=@"docs/help/questios.xlsx"'
```

- In the scheduled time, you can visit new blogs in the WordPress or X platform.
