//go:build js

package goji

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// customType is a type used to test marshalling
type customType struct {
	Name string `json:"name"`
	Age  int
}

func TestMarshalJS(t *testing.T) {
	expect := customType{
		Name: "Alice",
		Age:  42,
	}

	// marshal into a js.Value
	value, err := MarshalJS(expect)
	require.NoError(t, err)

	// value should be preserved and json tags respected
	assert.Equal(t, expect.Name, value.Get("name").String())
	assert.Equal(t, expect.Age, value.Get("Age").Int())
}

func TestUnmarshalJS(t *testing.T) {
	expect := customType{
		Name: "Bob",
		Age:  41,
	}

	// marshal into a js.Value
	value, err := MarshalJS(expect)
	require.NoError(t, err)

	// unmarshal back into a struct
	var actual customType
	err = UnmarshalJS(value, &actual)
	require.NoError(t, err)

	// should be the same
	assert.Equal(t, expect, actual)
}
