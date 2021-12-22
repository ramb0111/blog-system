package article

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Validatable interface {
	Validate() error
}

func payloadValidation(c *gin.Context, payload Validatable) error {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			// TODO: can error messgae be different?
			ErrResponse(http.StatusBadRequest, err),
		)
		return err
	}
	err = json.Unmarshal(bytes, &payload)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			ErrResponse(http.StatusBadRequest, err),
		)
		return err
	}

	err = payload.Validate()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			ErrResponse(http.StatusBadRequest, err),
		)
		return err
	}
	return err
}
