package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/ramb0111/blog-system/internal/api"
	"github.com/ramb0111/blog-system/internal/repository"
	"github.com/ramb0111/blog-system/pkg/server"
)

const (
	DYNAMO_DB_ENDPOINT = "http://localhost:8000"
)

func main() {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Endpoint: aws.String(DYNAMO_DB_ENDPOINT)})
	svc := server.NewServer(api.NewHandler(repository.NewRepository(db)), "")
	svc.Serve()
}
