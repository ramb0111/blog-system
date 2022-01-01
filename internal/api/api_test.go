package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	handlers := NewHandler(nil)
	assert.NotNil(t, handlers)
}
