//go:build js

package indexed_db

import (
	"syscall/js"
	"testing"

	"github.com/sourcenetwork/goji"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	var req RequestValue[DatabaseValue]
	upgradeNeeded := goji.EventListener(func(event goji.EventValue) {
		store := req.Result().CreateObjectStore("authors")
		store.CreateIndex("authors_age", js.ValueOf("age"), WithUnique(true))
	})
	defer upgradeNeeded.Release()

	req = Open(t.Name(), 1)
	req.EventTarget().AddEventListener(UpgradeNeededEvent, upgradeNeeded.Value)

	db, err := Await(req)
	require.NoError(t, err)
	defer db.Close()

	key := js.ValueOf(1)
	val := js.ValueOf(map[string]any{"name": "bob", "age": 65})

	transaction := db.Transaction(js.ValueOf("authors"), TransactionModeReadWrite)
	defer transaction.Abort()

	store := transaction.ObjectStore("authors")
	index := store.Index("authors_age")

	_, err = Await(store.Add(val, key))
	require.NoError(t, err)

	assert.Equal(t, "authors_age", index.Name())
	assert.Equal(t, "age", index.KeyPath().String())
	assert.Equal(t, false, index.MultiEntry())
	assert.Equal(t, true, index.Unique())

	count, err := Await(index.Count(js.ValueOf(65)))
	require.NoError(t, err)
	assert.Equal(t, 1, count.Int())

	actualVal, err := Await(index.Get(js.ValueOf(65)))
	require.NoError(t, err)
	assert.Equal(t, 65, actualVal.Get("age").Int())
	assert.Equal(t, "bob", actualVal.Get("name").String())
}
