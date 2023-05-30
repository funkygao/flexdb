package email

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestEvaluateCell(t *testing.T) {
	c := email{}
	v, err := c.EvaluateCell("foo@163.com", nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, "foo@163.com", v)
}
