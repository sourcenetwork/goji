//go:build js

package indexed_db

import (
	"testing"

	"github.com/sourcenetwork/goji"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseCreateObjectStore(t *testing.T) {
	var req RequestValue[DatabaseValue]
	upgradeNeeded := goji.EventListener(func(event goji.EventValue) {
		req.Result().CreateObjectStore("authors")
		req.Result().CreateObjectStore("books")
	})
	defer upgradeNeeded.Release()

	req = Open(t.Name(), 1)
	req.EventTarget().AddEventListener(UpgradeNeededEvent, upgradeNeeded.Value)

	db, err := Await(req)
	require.NoError(t, err)
	defer db.Close()

	names := db.ObjectStoreNames()
	assert.Equal(t, 2, names.Get("length").Int())
	assert.Equal(t, "authors", names.Index(0).String())
	assert.Equal(t, "books", names.Index(1).String())
}

func TestDatabaseDeleteObjectStore(t *testing.T) {
	var req RequestValue[DatabaseValue]
	upgradeNeeded := goji.EventListener(func(event goji.EventValue) {
		req.Result().CreateObjectStore("authors")
		req.Result().CreateObjectStore("books")
		req.Result().DeleteObjectStore("authors")
	})
	defer upgradeNeeded.Release()

	req = Open(t.Name(), 1)
	req.EventTarget().AddEventListener(UpgradeNeededEvent, upgradeNeeded.Value)

	db, err := Await(req)
	require.NoError(t, err)
	defer db.Close()

	names := db.ObjectStoreNames()
	assert.Equal(t, 1, names.Get("length").Int())
	assert.Equal(t, "books", names.Index(0).String())
}
