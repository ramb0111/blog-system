package article

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AddArticleRequestDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (dto *AddArticleRequestDTO) Validate() error {
	return validation.ValidateStruct(dto,
		validation.Field(&dto.Title, validation.Required),
		validation.Field(&dto.Content, validation.Required),
		validation.Field(&dto.Author, validation.Required),
	)
}

type addArticleRepository interface {
	AddArticle(context.Context, AddArticleRequestDTO) (string, error)
}

func AddArticleHandler(articleRepo addArticleRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := AddArticleRequestDTO{}
		if err := payloadValidation(c, &payload); err != nil {
			return
		}

		articleID, err := articleRepo.AddArticle(c, payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrResponse(http.StatusInternalServerError, err))
			return
		}

		articleResp := struct {
			ID string `json:"id"`
		}{ID: articleID}

		c.JSON(http.StatusCreated, SuccessResponse(http.StatusCreated, articleResp))
	}
}
