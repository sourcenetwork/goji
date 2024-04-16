//go:build js

package goji

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPromiseNew(t *testing.T) {
	var (
		res      []js.Value
		executor js.Func
	)

	executor = js.FuncOf(func(this js.Value, args []js.Value) any {
		executor.Release()
		res = args
		return js.Undefined()
	})

	value := Promise.New(executor)
	assert.NotNil(t, value)

	require.Len(t, res, 2)
	assert.Equal(t, js.TypeFunction, res[0].Type())
	assert.Equal(t, js.TypeFunction, res[1].Type())
}

func TestPromiseResolveThenAwait(t *testing.T) {
	value := js.ValueOf(1)
	prom := Promise.Resolve(value)

	res, err := Await(prom)
	require.NoError(t, err)

	require.Len(t, res, 1)
	assert.Equal(t, value, res[0])
}

func TestPromiseRejectThenAwait(t *testing.T) {
	value := Error.New("test")
	prom := Promise.Reject(js.Value(value))

	res, err := Await(prom)
	assert.Len(t, res, 0)
	assert.Equal(t, value, err)
}

func TestPromiseOfThenAwaitResolve(t *testing.T) {
	value := js.ValueOf(true)
	prom := PromiseOf(func(resolve, reject func(value js.Value)) {
		resolve(value)
	})

	res, err := Await(prom)
	require.NoError(t, err)

	require.Len(t, res, 1)
	assert.Equal(t, value, res[0])
}

func TestPromiseOfThenAwaitReject(t *testing.T) {
	value := Error.New("test reject")
	prom := PromiseOf(func(resolve, reject func(value js.Value)) {
		reject(js.Value(value))
	})

	res, err := Await(prom)
	assert.Len(t, res, 0)
	assert.Equal(t, value, err)
}
