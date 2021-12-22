package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramb0111/blog-system/internal/api/article"
)

type ArticleRepository interface {
	AddArticle(article article.AddArticleRequestDTO) (int, error)
	GetArticles() ([]article.GetArticleReponse, error)
	GetArticleByID(int) (article.GetArticleReponse, error)
}

func NewHandler(articleRepo ArticleRepository) http.Handler {
	engine := gin.Default()
	engine.GET("/articles/:id", article.GetArticlesByIDHandler(articleRepo))
	engine.GET("/articles", article.GetArticlesHandler(articleRepo))
	engine.POST("/articles", article.AddArticleHandler(articleRepo))
	return engine
}
