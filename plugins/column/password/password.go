package password

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*password)(nil)
	_ entity.Indexable    = (*password)(nil)
)

type password struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *password) ValidateCell(val string) error {
	return nil
}

func (password) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *password) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (password) IndexKind() entity.Index {
	return &entity.IndexStr{}
}
