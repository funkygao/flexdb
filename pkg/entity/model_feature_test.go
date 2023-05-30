package entity

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestModelFeature(t *testing.T) {
	var f ModelFeature
	assert.Equal(t, false, f.ChangeAuditEnabled())

	var b = [featureSize]byte{
		'1', '0', '1',
	}

	f, err := f.parseFeature(string(b[:]))
	assert.Equal(t, nil, err)
	assert.Equal(t, true, f.ChangeAuditEnabled())

	err = f.Scan(string(b[:]))
	assert.Equal(t, nil, err)

	val, err := f.Value()
	assert.Equal(t, nil, err)
	t.Logf("%v", val)
}
