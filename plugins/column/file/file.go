package file

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*file)(nil)
	// file cannot be indexed
)

type file struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *file) ValidateCell(val string) error {
	return nil
}

func (file) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}
