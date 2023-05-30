package qrcode

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*qrcode)(nil)
)

type qrcode struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *qrcode) ValidateCell(val string) error {
	return nil
}

func (qrcode) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}
