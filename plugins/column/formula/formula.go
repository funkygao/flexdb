package formula

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*formula)(nil)
)

type formula struct {
	*entity.Column
	ctx entity.PluginContext
}

func (f *formula) ValidateCell(val string) error {
	return nil
}

func (f *formula) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	_ = f.BizRule
	return val, nil
}
