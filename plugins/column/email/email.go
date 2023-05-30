package email

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*email)(nil)
	_ entity.Indexable    = (*email)(nil)
)

type email struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *email) ValidateCell(val string) error {
	return nil
}

func (email) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *email) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (email) IndexKind() entity.Index {
	return &entity.IndexStr{}
}
