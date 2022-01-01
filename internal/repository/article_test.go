package repository

import (
	"context"
	"errors"
	"sort"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/guregu/dynamo"
	"github.com/gusaul/go-dynamock"
	"github.com/ramb0111/blog-system/internal/api/article"
	"github.com/stretchr/testify/assert"
)

func mockRun(err error) CreateTableI {
	return &CreateTableIMock{
		RunFunc: func() error {
			return err
		},
	}
}

func TestNewRepository_Ok(t *testing.T) {
	testDB := dynamo.New(session.Must(session.NewSession()), &aws.Config{})

	createTableMock := func(name string, from interface{}) CreateTableI {
		return mockRun(nil)
	}
	repo := NewRepository(testDB, createTableMock)

	assert.Equal(t, ArticleTableName, repo.table.Name())
}

func TestNewRepository_Panic(t *testing.T) {
	testDB := dynamo.New(session.Must(session.NewSession()), &aws.Config{})

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	createTableMock := func(name string, from interface{}) CreateTableI {
		return mockRun(errors.New("create table error"))
	}
	NewRepository(testDB, createTableMock)
}

func getRepository() (*Repository, *dynamock.DynaMock) {
	mockDB, mock := dynamock.New()

	db := dynamo.NewFromIface(mockDB)

	return &Repository{
		table: db.Table("test-table"),
	}, mock
}

func TestAddArticle_Ok(t *testing.T) {
	repo, mock := getRepository()
	mock.ExpectPutItem().ToTable("test-table")

	_, err := repo.AddArticle(context.Background(), article.AddArticleRequestDTO{
		Title:   "some-title",
		Content: "some-content",
		Author:  "some-author",
	})

	assert.Nil(t, err)
}

func TestGetArticleByID_Ok(t *testing.T) {
	repo, mock := getRepository()

	dao := ArticleDAO{
		ID:      "9a48f420-2b5f-439a-8199-5fc9e8486e38",
		Title:   "some-title",
		Content: "some-content",
		Author:  "some-author",
	}
	item, err := dynamodbattribute.MarshalMap(dao)
	assert.Nil(t, err)

	mock.ExpectGetItem().ToTable("test-table").WillReturns(dynamodb.GetItemOutput{
		Item: item,
	})

	resp, err := repo.GetArticleByID(context.Background(), "9a48f420-2b5f-439a-8199-5fc9e8486e38")

	assert.Nil(t, err)
	expectedResp := article.GetArticleResponse{
		ID:      "9a48f420-2b5f-439a-8199-5fc9e8486e38",
		Title:   "some-title",
		Content: "some-content",
		Author:  "some-author",
	}
	assert.Equal(t, expectedResp, resp)
}

func TestGetArticles_Ok(t *testing.T) {
	repo, mock := getRepository()

	dao := []ArticleDAO{
		{
			ID:      "1a48f420-2b5f-439a-8199-5fc9e8486e46",
			Title:   "some-title",
			Content: "some-content",
			Author:  "some-author",
		},
		{
			ID:      "9a48f420-2b5f-439a-8199-5fc9e8486e38",
			Title:   "some-title",
			Content: "some-content",
			Author:  "some-author",
		},
	}
	item1, err := dynamodbattribute.MarshalMap(dao[0])
	assert.Nil(t, err)

	item2, err := dynamodbattribute.MarshalMap(dao[1])
	assert.Nil(t, err)

	items := []map[string]*dynamodb.AttributeValue{item1, item2}

	mock.ExpectScan().Table("test-table").WillReturns(dynamodb.ScanOutput{
		Items: items,
	})

	resp, err := repo.GetArticles(context.Background())

	assert.Nil(t, err)
	expectedResp := []article.GetArticleResponse{
		{
			ID:      "1a48f420-2b5f-439a-8199-5fc9e8486e46",
			Title:   "some-title",
			Content: "some-content",
			Author:  "some-author",
		},
		{
			ID:      "9a48f420-2b5f-439a-8199-5fc9e8486e38",
			Title:   "some-title",
			Content: "some-content",
			Author:  "some-author",
		},
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].ID < resp[j].ID
	})
	sort.Slice(expectedResp, func(i, j int) bool {
		return expectedResp[i].ID < expectedResp[j].ID
	})

	assert.Equal(t, expectedResp, resp)
}
