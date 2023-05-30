package integer

import (
	"strconv"

	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin = (*integer)(nil)
	_ entity.Indexable    = (*integer)(nil)
)

type integer struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *integer) ValidateCell(val string) error {
	return nil
}

func (integer) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return strconv.ParseInt(val, 10, 64)
}

func (c *integer) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	idx := &entity.IndexInt{Slot: c.Slot}
	v, err := c.EvaluateCell(val, nil)
	if err != nil {
		return nil, err
	}

	idx.Val = v.(int64)
	return idx, nil
}

func (integer) IndexKind() entity.Index {
	return &entity.IndexInt{}
}
