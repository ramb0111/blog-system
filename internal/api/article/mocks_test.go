// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package article

import (
	"context"
	"sync"
)

// Ensure, that addArticleRepositoryMock does implement addArticleRepository.
// If this is not the case, regenerate this file with moq.
var _ addArticleRepository = &addArticleRepositoryMock{}

// addArticleRepositoryMock is a mock implementation of addArticleRepository.
//
// 	func TestSomethingThatUsesaddArticleRepository(t *testing.T) {
//
// 		// make and configure a mocked addArticleRepository
// 		mockedaddArticleRepository := &addArticleRepositoryMock{
// 			AddArticleFunc: func(contextMoqParam context.Context, addArticleRequestDTO AddArticleRequestDTO) (string, error) {
// 				panic("mock out the AddArticle method")
// 			},
// 		}
//
// 		// use mockedaddArticleRepository in code that requires addArticleRepository
// 		// and then make assertions.
//
// 	}
type addArticleRepositoryMock struct {
	// AddArticleFunc mocks the AddArticle method.
	AddArticleFunc func(contextMoqParam context.Context, addArticleRequestDTO AddArticleRequestDTO) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddArticle holds details about calls to the AddArticle method.
		AddArticle []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// AddArticleRequestDTO is the addArticleRequestDTO argument value.
			AddArticleRequestDTO AddArticleRequestDTO
		}
	}
	lockAddArticle sync.RWMutex
}

// AddArticle calls AddArticleFunc.
func (mock *addArticleRepositoryMock) AddArticle(contextMoqParam context.Context, addArticleRequestDTO AddArticleRequestDTO) (string, error) {
	if mock.AddArticleFunc == nil {
		panic("addArticleRepositoryMock.AddArticleFunc: method is nil but addArticleRepository.AddArticle was just called")
	}
	callInfo := struct {
		ContextMoqParam      context.Context
		AddArticleRequestDTO AddArticleRequestDTO
	}{
		ContextMoqParam:      contextMoqParam,
		AddArticleRequestDTO: addArticleRequestDTO,
	}
	mock.lockAddArticle.Lock()
	mock.calls.AddArticle = append(mock.calls.AddArticle, callInfo)
	mock.lockAddArticle.Unlock()
	return mock.AddArticleFunc(contextMoqParam, addArticleRequestDTO)
}

// AddArticleCalls gets all the calls that were made to AddArticle.
// Check the length with:
//     len(mockedaddArticleRepository.AddArticleCalls())
func (mock *addArticleRepositoryMock) AddArticleCalls() []struct {
	ContextMoqParam      context.Context
	AddArticleRequestDTO AddArticleRequestDTO
} {
	var calls []struct {
		ContextMoqParam      context.Context
		AddArticleRequestDTO AddArticleRequestDTO
	}
	mock.lockAddArticle.RLock()
	calls = mock.calls.AddArticle
	mock.lockAddArticle.RUnlock()
	return calls
}

// Ensure, that getArticlesRepositoryMock does implement getArticlesRepository.
// If this is not the case, regenerate this file with moq.
var _ getArticlesRepository = &getArticlesRepositoryMock{}

// getArticlesRepositoryMock is a mock implementation of getArticlesRepository.
//
// 	func TestSomethingThatUsesgetArticlesRepository(t *testing.T) {
//
// 		// make and configure a mocked getArticlesRepository
// 		mockedgetArticlesRepository := &getArticlesRepositoryMock{
// 			GetArticlesFunc: func(contextMoqParam context.Context) ([]GetArticleResponse, error) {
// 				panic("mock out the GetArticles method")
// 			},
// 		}
//
// 		// use mockedgetArticlesRepository in code that requires getArticlesRepository
// 		// and then make assertions.
//
// 	}
type getArticlesRepositoryMock struct {
	// GetArticlesFunc mocks the GetArticles method.
	GetArticlesFunc func(contextMoqParam context.Context) ([]GetArticleResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetArticles holds details about calls to the GetArticles method.
		GetArticles []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
	}
	lockGetArticles sync.RWMutex
}

// GetArticles calls GetArticlesFunc.
func (mock *getArticlesRepositoryMock) GetArticles(contextMoqParam context.Context) ([]GetArticleResponse, error) {
	if mock.GetArticlesFunc == nil {
		panic("getArticlesRepositoryMock.GetArticlesFunc: method is nil but getArticlesRepository.GetArticles was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockGetArticles.Lock()
	mock.calls.GetArticles = append(mock.calls.GetArticles, callInfo)
	mock.lockGetArticles.Unlock()
	return mock.GetArticlesFunc(contextMoqParam)
}

// GetArticlesCalls gets all the calls that were made to GetArticles.
// Check the length with:
//     len(mockedgetArticlesRepository.GetArticlesCalls())
func (mock *getArticlesRepositoryMock) GetArticlesCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockGetArticles.RLock()
	calls = mock.calls.GetArticles
	mock.lockGetArticles.RUnlock()
	return calls
}

// Ensure, that getArticleRepositoryMock does implement getArticleRepository.
// If this is not the case, regenerate this file with moq.
var _ getArticleRepository = &getArticleRepositoryMock{}

// getArticleRepositoryMock is a mock implementation of getArticleRepository.
//
// 	func TestSomethingThatUsesgetArticleRepository(t *testing.T) {
//
// 		// make and configure a mocked getArticleRepository
// 		mockedgetArticleRepository := &getArticleRepositoryMock{
// 			GetArticleByIDFunc: func(contextMoqParam context.Context, s string) (GetArticleResponse, error) {
// 				panic("mock out the GetArticleByID method")
// 			},
// 		}
//
// 		// use mockedgetArticleRepository in code that requires getArticleRepository
// 		// and then make assertions.
//
// 	}
type getArticleRepositoryMock struct {
	// GetArticleByIDFunc mocks the GetArticleByID method.
	GetArticleByIDFunc func(contextMoqParam context.Context, s string) (GetArticleResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetArticleByID holds details about calls to the GetArticleByID method.
		GetArticleByID []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
		}
	}
	lockGetArticleByID sync.RWMutex
}

// GetArticleByID calls GetArticleByIDFunc.
func (mock *getArticleRepositoryMock) GetArticleByID(contextMoqParam context.Context, s string) (GetArticleResponse, error) {
	if mock.GetArticleByIDFunc == nil {
		panic("getArticleRepositoryMock.GetArticleByIDFunc: method is nil but getArticleRepository.GetArticleByID was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
	}
	mock.lockGetArticleByID.Lock()
	mock.calls.GetArticleByID = append(mock.calls.GetArticleByID, callInfo)
	mock.lockGetArticleByID.Unlock()
	return mock.GetArticleByIDFunc(contextMoqParam, s)
}

// GetArticleByIDCalls gets all the calls that were made to GetArticleByID.
// Check the length with:
//     len(mockedgetArticleRepository.GetArticleByIDCalls())
func (mock *getArticleRepositoryMock) GetArticleByIDCalls() []struct {
	ContextMoqParam context.Context
	S               string
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
	}
	mock.lockGetArticleByID.RLock()
	calls = mock.calls.GetArticleByID
	mock.lockGetArticleByID.RUnlock()
	return calls
}
