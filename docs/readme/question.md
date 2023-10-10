# REST API Question

## list questions

### GET ```127.0.0.1:4030/questions```
### Response
```json
{
    "results": [
      {
        "id": "4",
        "created_at": "2023-05-12T11:12:04.713429Z",
        "question": "Unauthorized access to ICS devices: Blockchain can enable access control mechanisms that can allow only authorized personnel to access the system.",
        "rule": "write a professional blog post with maximum 500 word, and a summary of the post for linkedin and twitter for the above topic and return the post with the below structure post title: write title here write the remaining article without subtitles linkedin post: write the summary here twitter post: write the summary here",
        "status": "0"
      },
      {
        "id": "5",
        "created_at": "2023-05-12T11:12:04.713429Z",
        "question": "Unauthorized access to ICS devices: Blockchain can enable access control mechanisms that can allow only authorized personnel to access the system.",
        "rule": "write a professional blog post with maximum 500 word, and a summary of the post for linkedin and twitter for the above topic and return the post with the below structure post title: write title here write the remaining article without subtitles linkedin post: write the summary here twitter post: write the summary here",
        "status": "1"
      }
    ]
}
```

### GET ```127.0.0.1:4030/questions?status=0```
### Response
```json
{
    "results": [
      {
        "id": "4",
        "created_at": "2023-05-12T11:12:04.713429Z",
        "question": "Unauthorized access to ICS devices: Blockchain can enable access control mechanisms that can allow only authorized personnel to access the system.",
        "rule": "write a professional blog post with maximum 500 word, and a summary of the post for linkedin and twitter for the above topic and return the post with the below structure post title: write title here write the remaining article without subtitles linkedin post: write the summary here twitter post: write the summary here",
        "status": "0"
      }
    ]
}
```

### GET ```127.0.0.1:4030/questions?status=1```
### Response
```json
{
    "results": [
      {
        "id": "4",
        "created_at": "2023-05-12T11:12:04.713429Z",
        "question": "Unauthorized access to ICS devices: Blockchain can enable access control mechanisms that can allow only authorized personnel to access the system.",
        "rule": "write a professional blog post with maximum 500 word, and a summary of the post for linkedin and twitter for the above topic and return the post with the below structure post title: write title here write the remaining article without subtitles linkedin post: write the summary here twitter post: write the summary here",
        "status": "1"
      }
    ]
}
```
## upload questions
```http request
POST {{baseUrl}}/question?tag=1,4,5&category=2,5
Content-Type: multipart/form-data;
Content-Disposition: form-data; name="file.xlsx";
```
