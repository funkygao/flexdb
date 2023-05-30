package url

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*url)(nil)
	_ entity.Indexable    = (*url)(nil)
)

type url struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *url) ValidateCell(val string) error {
	return nil
}

func (url) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *url) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (url) IndexKind() entity.Index {
	return &entity.IndexStr{}
}
