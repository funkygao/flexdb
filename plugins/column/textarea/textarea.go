package textarea

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*textarea)(nil)
	// textarea cannot be indexed TODO fulltext search in ES
)

type textarea struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *textarea) ValidateCell(val string) error {
	return nil
}

func (textarea) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}
