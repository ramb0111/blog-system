package article

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetArticleReponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type getArticleRepository interface {
	GetArticles() ([]GetArticleReponse, error)
	GetArticleByID(int) (GetArticleReponse, error)
}

func GetArticlesHandler(articleRepo getArticleRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		articles, err := articleRepo.GetArticles()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrResponse(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusCreated, SuccessResponse(http.StatusCreated, articles))
	}
}

func GetArticlesByIDHandler(articleRepo getArticleRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrResponse(http.StatusBadRequest, err))
			return
		}
		article, err := articleRepo.GetArticleByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrResponse(http.StatusInternalServerError, err))
			return
		}

		articles := []GetArticleReponse{article}
		c.JSON(http.StatusCreated, SuccessResponse(http.StatusCreated, articles))
	}
}
