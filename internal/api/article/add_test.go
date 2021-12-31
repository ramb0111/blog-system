package article

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tcs := []struct {
		dto         *AddArticleRequestDTO
		expectedErr string
	}{
		{
			dto:         &AddArticleRequestDTO{},
			expectedErr: "author: cannot be blank; content: cannot be blank; title: cannot be blank.",
		},
		{
			dto: &AddArticleRequestDTO{
				Content: "some-content",
			},
			expectedErr: "author: cannot be blank; title: cannot be blank.",
		},
		{
			dto: &AddArticleRequestDTO{
				Title:  "some-title",
				Author: "some-author",
			},
			expectedErr: "content: cannot be blank.",
		},
		{
			dto: &AddArticleRequestDTO{
				Author: "some-author",
			},
			expectedErr: "content: cannot be blank; title: cannot be blank.",
		},
		{
			dto: &AddArticleRequestDTO{
				Title:   "some-title",
				Content: "some-content",
				Author:  "some-author",
			},
		},
	}

	for _, tc := range tcs {
		if tc.expectedErr == "" {
			assert.Nil(t, tc.dto.Validate())
		} else {
			assert.EqualError(t, tc.dto.Validate(), tc.expectedErr)
		}
	}
}

func mockaddArticle(ID string, err error) addArticleRepository {
	return &addArticleRepositoryMock{
		AddArticleFunc: func(contextMoqParam context.Context, addArticleRequestDTO AddArticleRequestDTO) (string, error) {
			return ID, err
		},
	}
}

func TestAddArticleHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tcs := []struct {
		mockedID  string
		mockedErr error
		dto       interface{}

		expectedBody string
		expectedCode int
	}{
		{
			dto: AddArticleRequestDTO{
				Title: "some-title",
			},
			expectedBody: `{"status":400,"message":"author: cannot be blank; content: cannot be blank.","data":null}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			dto: AddArticleRequestDTO{
				Title:  "some-title",
				Author: "some-author",
			},
			expectedBody: `{"status":400,"message":"content: cannot be blank.","data":null}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			dto: AddArticleRequestDTO{
				Title:   "some-title",
				Author:  "some-author",
				Content: "some-content",
			},
			expectedBody: `{"status":500,"message":"some-db-error","data":null}`,
			expectedCode: http.StatusInternalServerError,
			mockedErr:    errors.New("some-db-error"),
		},
		{
			dto: AddArticleRequestDTO{
				Title:   "some-title",
				Author:  "some-author",
				Content: "some-content",
			},
			expectedBody: `{"status":201,"message":"Success","data":{"id":"9a48f420-2b5f-439a-8199-5fc9e8486e38"}}`,
			expectedCode: http.StatusCreated,
			mockedID:     "9a48f420-2b5f-439a-8199-5fc9e8486e38",
		},
	}

	for _, tc := range tcs {
		r := gin.Default()
		r.POST("/articles", AddArticleHandler(mockaddArticle(tc.mockedID, tc.mockedErr)))

		payload, err := json.Marshal(tc.dto)
		assert.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(payload))
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		resp, err := ioutil.ReadAll(w.Body)
		assert.Nil(t, err)
		assert.Equal(t, tc.expectedBody, string(resp))
		assert.Equal(t, tc.expectedCode, w.Code)
	}

}
