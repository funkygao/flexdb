package rating

import (
	"strconv"

	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*rating)(nil)
)

type rating struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *rating) ValidateCell(val string) error {
	_, err := c.EvaluateCell(val, nil)
	return err
}

func (rating) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return strconv.ParseInt(val, 10, 64)
}
