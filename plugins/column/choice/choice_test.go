package choice

import (
	"testing"

	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/funkygao/assert"
)

func TestEvaluateCell(t *testing.T) {
	c := choice{Column: &entity.Column{}}
	v, err := c.EvaluateCell("5", nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(5), v)
}
