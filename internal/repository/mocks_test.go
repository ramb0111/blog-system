// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"sync"
)

// Ensure, that CreateTableIMock does implement CreateTableI.
// If this is not the case, regenerate this file with moq.
var _ CreateTableI = &CreateTableIMock{}

// CreateTableIMock is a mock implementation of CreateTableI.
//
// 	func TestSomethingThatUsesCreateTableI(t *testing.T) {
//
// 		// make and configure a mocked CreateTableI
// 		mockedCreateTableI := &CreateTableIMock{
// 			RunFunc: func() error {
// 				panic("mock out the Run method")
// 			},
// 		}
//
// 		// use mockedCreateTableI in code that requires CreateTableI
// 		// and then make assertions.
//
// 	}
type CreateTableIMock struct {
	// RunFunc mocks the Run method.
	RunFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Run holds details about calls to the Run method.
		Run []struct {
		}
	}
	lockRun sync.RWMutex
}

// Run calls RunFunc.
func (mock *CreateTableIMock) Run() error {
	if mock.RunFunc == nil {
		panic("CreateTableIMock.RunFunc: method is nil but CreateTableI.Run was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRun.Lock()
	mock.calls.Run = append(mock.calls.Run, callInfo)
	mock.lockRun.Unlock()
	return mock.RunFunc()
}

// RunCalls gets all the calls that were made to Run.
// Check the length with:
//     len(mockedCreateTableI.RunCalls())
func (mock *CreateTableIMock) RunCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRun.RLock()
	calls = mock.calls.Run
	mock.lockRun.RUnlock()
	return calls
}