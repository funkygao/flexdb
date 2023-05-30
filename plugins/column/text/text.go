package text

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*text)(nil)
	_ entity.Indexable    = (*text)(nil)
)

type text struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *text) ValidateCell(val string) error {
	return nil
}

func (text) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *text) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (text) IndexKind() entity.Index {
	return &entity.IndexStr{}
}
