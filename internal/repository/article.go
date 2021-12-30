package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	api_article "github.com/ramb0111/blog-system/internal/api/article"
)

const (
	ArticleTableName = "article"
)

type articleDAO struct {
	ID string

	Title   string
	Content string
	Author  string
}

type Repository struct {
	db    *dynamo.DB
	table dynamo.Table

	now func() time.Time
}

func NewRepository(db *dynamo.DB) *Repository {
	return &Repository{
		db:    db,
		table: db.Table(ArticleTableName),
		now:   time.Now,
	}
}

func (r *Repository) AddArticle(ctx context.Context, article api_article.AddArticleRequestDTO) (string, error) {
	articleDAO := articleDAO{
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

func (r *Repository) GetArticle(ctx context.Context, ID string) (*api_article.GetArticleResponse, error) {
	var articleDAO articleDAO
	if err := r.table.Get("ID", ID).OneWithContext(ctx, &articleDAO); err != nil {
		return nil, err
	}

	return &api_article.GetArticleResponse{
		ID:      articleDAO.ID,
		Author:  articleDAO.Author,
		Title:   articleDAO.Title,
		Content: articleDAO.Content,
	}, nil
}

func (r *Repository) GetArticles(ctx context.Context) (*[]api_article.GetArticleResponse, error) {
	var articlesDAO []articleDAO
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

	return &articlesDTO, nil
}
