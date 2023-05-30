package dto

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestCriteria(t *testing.T) {
	c := Criteria{
		{Key: "name", Op: "=", Val: "funky"},
		{Key: "birthdate", Op: ">=", Val: "1986-11-02 15:04:05"},
	}
	assert.Equal(t, true, c.Valid())

	assert.Equal(t, 2, c.Size())
	c = c.Append(CriteriaItem{Key: "name1", Op: "=", Val: "funky"})
	assert.Equal(t, 3, c.Size())

	var c1 Criteria
	assert.Equal(t, 0, c1.Size())
	c1 = c1.Append(CriteriaItem{Key: "name1", Op: "=", Val: "funky"})
	assert.Equal(t, 1, c1.Size())
}
