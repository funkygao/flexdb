package mobile

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*mobile)(nil)
	_ entity.Indexable    = (*mobile)(nil)
)

type mobile struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *mobile) ValidateCell(val string) error {
	return nil
}

func (mobile) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *mobile) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (mobile) IndexKind() entity.Index {
	return &entity.IndexStr{}
}
