package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/ramb0111/blog-system/internal/api"
	"github.com/ramb0111/blog-system/internal/repository"
	"github.com/ramb0111/blog-system/pkg/server"
)

var (
	DYNAMO_DB_ENDPOINT    = os.Getenv("DYNAMO_DB_ENDPOINT")
	AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_ACCESS_KEY_SECRET = os.Getenv("AWS_ACCESS_KEY_SECRET")
)

func main() {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess,
		aws.NewConfig().
			// To enable debugging
			// WithLogLevel(aws.LogDebugWithHTTPBody).
			WithEndpoint(DYNAMO_DB_ENDPOINT).
			WithRegion("us-east-1").
			WithCredentials(credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_ACCESS_KEY_SECRET, "")),
	)

	createTable := func(name string, from interface{}) repository.CreateTableI {
		return db.CreateTable(name, from)
	}
	server.NewServer(api.NewHandler(repository.NewRepository(db, createTable)), "").Serve()
}
