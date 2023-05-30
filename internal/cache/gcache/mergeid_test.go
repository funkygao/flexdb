package gcache

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestMergeID(t *testing.T) {
	id := int64(323232323)
	hintID := int64(5)
	mergeID := generateMergeID(hintID, id)
	h, i := parseMergeID(mergeID)
	assert.Equal(t, h, hintID)
	assert.Equal(t, i, id)
}
