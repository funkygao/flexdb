package city

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*city)(nil)
	_ entity.Indexable    = (*city)(nil)
)

type city struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *city) ValidateCell(val string) error {
	return nil
}

func (city) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *city) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (city) IndexKind() entity.Index {
	return &entity.IndexStr{}
}
