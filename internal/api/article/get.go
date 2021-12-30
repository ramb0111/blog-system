package article

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetArticleResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type getArticleRepository interface {
	GetArticles(context.Context) ([]GetArticleResponse, error)
	GetArticleByID(context.Context, string) (GetArticleResponse, error)
}

func GetArticlesHandler(articleRepo getArticleRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		articles, err := articleRepo.GetArticles(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrResponse(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, articles))
	}
}

func GetArticlesByIDHandler(articleRepo getArticleRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		article, err := articleRepo.GetArticleByID(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrResponse(http.StatusInternalServerError, err))
			return
		}

		articles := []GetArticleResponse{article}
		c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, articles))
	}
}
