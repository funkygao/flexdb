package mobile

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestEvaluateCell(t *testing.T) {
	c := mobile{}
	v, err := c.EvaluateCell("13910987654", nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, "13910987654", v)
}
