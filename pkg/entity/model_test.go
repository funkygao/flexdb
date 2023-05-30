package entity

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestModelDesc(t *testing.T) {
	m := Model{AppID: 1}
	m.Slots = []*Column{
		{Name: "a", Kind: ColumnAutoNumber},
	}
	assert.Equal(t, len(builtinColumns)+len(m.Slots), m.TotalColumnsN())
}
