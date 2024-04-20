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

	request := Open("test", 1)
	request.AddEventListener(UpgradeNeededEvent, upgradeNeeded.Value)

	res, err := Await(request)
	require.NoError(t, err)

	// upgradeNeeded should have been called by this point
	assert.True(t, upgradeCalled)

	db := DatabaseValue{goji.EventTargetValue(res)}
	assert.Equal(t, "test", db.Name())
}
