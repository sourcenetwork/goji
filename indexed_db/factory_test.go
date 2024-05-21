//go:build js

package indexed_db

import (
	"testing"

	"github.com/sourcenetwork/goji"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	upgradeCalled := false
	upgradeNeeded := goji.EventListener(func(event goji.EventValue) {
		upgradeCalled = true
	})
	defer upgradeNeeded.Release()

	request := Open(t.Name(), 1)
	request.EventTarget().AddEventListener(UpgradeNeededEvent, upgradeNeeded.Value)

	db, err := Await(request)
	require.NoError(t, err)
	defer db.Close()

	// upgradeNeeded should have been called by this point
	assert.True(t, upgradeCalled)
	assert.Equal(t, 1, db.Version())
	assert.Equal(t, t.Name(), db.Name())
}

func TestDelete(t *testing.T) {
	upgradeNeeded := goji.EventListener(func(event goji.EventValue) {
		// do nothing
	})
	defer upgradeNeeded.Release()

	openReq := Open(t.Name(), 1)
	openReq.EventTarget().AddEventListener(UpgradeNeededEvent, upgradeNeeded.Value)

	db, err := Await(openReq)
	require.NoError(t, err)
	db.Close() // must close before deleting

	deleteReq := DeleteDatabase(t.Name())
	_, err = Await(deleteReq)
	require.NoError(t, err)
}
