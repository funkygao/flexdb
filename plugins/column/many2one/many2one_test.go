package many2one

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestEvaluateCell(t *testing.T) {
	c := many2one{}
	v, err := c.EvaluateCell("5", nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(5), v)
}
