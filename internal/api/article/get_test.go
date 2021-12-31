package article

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockGetArticles(resp []GetArticleResponse, err error) getArticlesRepository {
	return &getArticlesRepositoryMock{
		GetArticlesFunc: func(contextMoqParam context.Context) ([]GetArticleResponse, error) {
			return resp, err
		},
	}
}

func TestGetArticlesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tcs := []struct {
		mockedResp []GetArticleResponse
		mockedErr  error

		expectedBody string
		expectedCode int
	}{
		{

			mockedResp: []GetArticleResponse{
				{
					Title:   "some-title",
					Content: "some-content",
					Author:  "some-author",
				},
				{
					Title:   "some-new-title",
					Content: "some-new-content",
					Author:  "some-new-author",
				},
			},
			expectedBody: `{"status":200,"message":"Success","data":[{"id":"","title":"some-title","content":"some-content","author":"some-author"},{"id":"","title":"some-new-title","content":"some-new-content","author":"some-new-author"}]}`,
			expectedCode: http.StatusOK,
		},
		{
			mockedErr:    errors.New("some-db-error"),
			expectedBody: `{"status":500,"message":"some-db-error","data":null}`,
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range tcs {
		r := gin.Default()
		r.GET("/articles", GetArticlesHandler(mockGetArticles(tc.mockedResp, tc.mockedErr)))

		req, err := http.NewRequest(http.MethodGet, "/articles", nil)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		resp, err := ioutil.ReadAll(w.Body)
		assert.Nil(t, err)
		assert.Equal(t, tc.expectedBody, string(resp))
		assert.Equal(t, tc.expectedCode, w.Code)
	}

}

func mockGetArticle(resp GetArticleResponse, err error) getArticleRepository {
	return &getArticleRepositoryMock{
		GetArticleByIDFunc: func(contextMoqParam context.Context, s string) (GetArticleResponse, error) {
			return resp, err
		},
	}
}

func TestGetArticlesByIDHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tcs := []struct {
		mockedResp GetArticleResponse
		mockedErr  error

		expectedBody string
		expectedCode int
	}{
		{

			mockedResp: GetArticleResponse{
				Title:   "some-title",
				Content: "some-content",
				Author:  "some-author",
			},
			expectedBody: `{"status":200,"message":"Success","data":[{"id":"","title":"some-title","content":"some-content","author":"some-author"}]}`,
			expectedCode: http.StatusOK,
		},
		{
			mockedErr:    errors.New("some-db-error"),
			expectedBody: `{"status":500,"message":"some-db-error","data":null}`,
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range tcs {
		r := gin.Default()
		r.GET("/articles", GetArticlesByIDHandler(mockGetArticle(tc.mockedResp, tc.mockedErr)))

		req, err := http.NewRequest(http.MethodGet, "/articles", nil)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		resp, err := ioutil.ReadAll(w.Body)
		assert.Nil(t, err)
		assert.Equal(t, tc.expectedBody, string(resp))
		assert.Equal(t, tc.expectedCode, w.Code)
	}

}
