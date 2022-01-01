package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	api_article "github.com/ramb0111/blog-system/internal/api/article"
)

const (
	ArticleTableName = "article"
)

type ArticleDAO struct {
	ID string `dynamo:"ID,hash"`

	Title   string
	Content string
	Author  string
}

type Repository struct {
	table dynamo.Table
}

type CreateTableI interface {
	Run() error
}

//To create a new article table for performing db operations
func NewRepository(db *dynamo.DB, createTable func(name string, from interface{}) CreateTableI) *Repository {
	if err := createTable(ArticleTableName, ArticleDAO{}).Run(); err != nil {
		panic(err)
	}

	return &Repository{
		table: db.Table(ArticleTableName),
	}
}

// AddArticle to add article instance into db
func (r *Repository) AddArticle(ctx context.Context, article api_article.AddArticleRequestDTO) (string, error) {
	articleDAO := ArticleDAO{
		ID:      uuid.New().String(),
		Author:  article.Author,
		Title:   article.Title,
		Content: article.Content,
	}

	if err := r.table.Put(&articleDAO).If("attribute_not_exists(ID)").RunWithContext(ctx); err != nil {
		return "", err
	}
	return articleDAO.ID, nil
}

// GetArticleByID to get article by given ID from db
func (r *Repository) GetArticleByID(ctx context.Context, ID string) (api_article.GetArticleResponse, error) {
	var articleDAO ArticleDAO
	if err := r.table.Get("ID", ID).OneWithContext(ctx, &articleDAO); err != nil {
		return api_article.GetArticleResponse{}, err
	}

	return api_article.GetArticleResponse{
		ID:      articleDAO.ID,
		Author:  articleDAO.Author,
		Title:   articleDAO.Title,
		Content: articleDAO.Content,
	}, nil
}

// GetArticles to get all the articles from db
func (r *Repository) GetArticles(ctx context.Context) ([]api_article.GetArticleResponse, error) {
	var articlesDAO []ArticleDAO
	if err := r.table.Scan().AllWithContext(ctx, &articlesDAO); err != nil {
		return nil, err
	}

	articlesDTO := []api_article.GetArticleResponse{}
	for _, article := range articlesDAO {
		articlesDTO = append(articlesDTO, api_article.GetArticleResponse{
			ID:      article.ID,
			Author:  article.Author,
			Title:   article.Title,
			Content: article.Content,
		})
	}

	return articlesDTO, nil
}
