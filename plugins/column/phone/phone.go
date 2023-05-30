package phone

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*phone)(nil)
	_ entity.Indexable    = (*phone)(nil)
)

type phone struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *phone) ValidateCell(val string) error {
	return nil
}

func (phone) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *phone) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (phone) IndexKind() entity.Index {
	return &entity.IndexStr{}
}
