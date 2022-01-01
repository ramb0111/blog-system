# Blog System
This blog system supports insertion and fetching of articles.


Project Layout - https://github.com/golang-standards/project-layout

## To Run the app
```
make run
```
| Note: make run will execute the test cases first

## Host to connect
```
localhost:8080
```


## To enable debugging 
uncomment this line `// WithLogLevel(aws.LogDebugWithHTTPBody).`



## Apis supported
### Create an article
- Method: `POST`
- Path: `/articles`
- Request Body:
```
{
    "title": "Hello World",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
    "author": "John",
}
```
- Response Header: `HTTP 201`
- Response Body:
```
{
    "status": 201,
    "message": "Success",
    "data": {
      "id": <article_id>
    }
}
```
or
- Response Header: `HTTP <HTTP_CODE>`
- Response Body:
```
{
    "status": <HTTP_CODE>,
    "message": <ERROR_DESCRIPTION>,
    "data": null
}
```

### Get article by id
- Method: `GET`
- Path: `/articles/<article_id>`
- Response Header: `HTTP 200`
- Response Body:
```
{
    "status": 200,
    "message": "Success",
    "data": [
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      }
    ]
}
```
or
- Response Header: `HTTP <HTTP_CODE>`
- Response Body:
```
{
    "status": <HTTP_CODE>,
    "message": <ERROR_DESCRIPTION>,
    "data": null
}
```

### Get all article
- Method: `GET`
- Path: `/articles`
- Response Header: `HTTP 200`
- Response Body:
```
{
    "status": 200,
    "message": "Success",
    "data": [
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      },
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      }
    ]
}
```
or
- Response Header: `HTTP <HTTP_CODE>`
- Response Body:
```
{
    "status": <HTTP_CODE>,
    "message": <ERROR_DESCRIPTION>,
    "data": null
}
```