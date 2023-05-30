package entity

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestUniqueIndexTableName(t *testing.T) {
	i := &IndexInt{}
	assert.Equal(t, "IndexInt", i.TableName())
	i.unique = true
	assert.Equal(t, "IndexIntUniq", i.TableName())

	j := &IndexStr{}
	assert.Equal(t, "IndexStr", j.TableName())
	j.unique = true
	assert.Equal(t, "IndexStrUniq", j.TableName())

	k := &IndexTime{}
	assert.Equal(t, "IndexTime", k.TableName())
	k.unique = true
	assert.Equal(t, "IndexTimeUniq", k.TableName())
}
