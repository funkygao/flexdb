package sequence

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestEvaluateCell(t *testing.T) {
	c := sequence{}
	v, err := c.EvaluateCell("5", nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(5), v)
}
