//go:build js

package goji

import (
	"fmt"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ (error) = (*CustomError)(nil)

type CustomError struct{}

func (CustomError) Error() string {
	return "custom error message"
}

func TestErrorNew(t *testing.T) {
	err := Error.New("test message")
	assert.Equal(t, "Error: test message", err.Error())
}

func TestWrapError(t *testing.T) {
	err := fmt.Errorf("something went wrong")
	wrapped := WrapError(err)
	assert.Equal(t, fmt.Sprintf("errorString: %s", err.Error()), wrapped.Error())
	assert.Equal(t, "errorString", js.Value(wrapped).Get("name").String())
}

func TestWrapErrorWithCustomError(t *testing.T) {
	err := &CustomError{}
	wrapped := WrapError(err)
	assert.Equal(t, fmt.Sprintf("CustomError: %s", err.Error()), wrapped.Error())
	assert.Equal(t, "CustomError", js.Value(wrapped).Get("name").String())
}
