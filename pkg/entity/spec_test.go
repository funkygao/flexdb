package entity

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestBuiltinColumns(t *testing.T) {
	Prepare()

	cols := []string{
		"id",
		"modelID",
		"model_id",
		"orgID",
		"org_id",
		"mtime",
		"ctime",
		"ts",
		"deleteD",
		"slug",
		"SLUG",
	}

	for _, c := range cols {
		assert.Equal(t, true, columnNameReserved(c))
	}

	// defensive test: if accidentally added a builtin column, test will fail
	assert.Equal(t, 6, len(builtinColumns))
}
