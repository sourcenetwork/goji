//go:build js

package goji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNew(t *testing.T) {
	err := Error.New("test message")
	assert.Equal(t, "Error: test message", err.Error())
}
