package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramb0111/blog-system/internal/api/article"
)

type ArticleRepository interface {
	AddArticle(context.Context, article.AddArticleRequestDTO) (string, error)
	GetArticles(context.Context) ([]article.GetArticleResponse, error)
	GetArticleByID(context.Context, string) (article.GetArticleResponse, error)
}

func NewHandler(articleRepo ArticleRepository) http.Handler {
	engine := gin.Default()
	engine.GET("/articles/:id", article.GetArticlesByIDHandler(articleRepo))
	engine.GET("/articles", article.GetArticlesHandler(articleRepo))
	engine.POST("/articles", article.AddArticleHandler(articleRepo))
	return engine
}
