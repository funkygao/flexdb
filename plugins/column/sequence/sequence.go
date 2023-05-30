package sequence

import (
	"strconv"

	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin       = (*sequence)(nil)
	_ entity.CellValueGenerator = (*sequence)(nil)
	_ entity.Indexable          = (*sequence)(nil)
)

type sequence struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *sequence) ValidateCell(val string) error {
	return nil
}

func (sequence) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return strconv.ParseInt(val, 10, 64)
}

func (c *sequence) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	// sequence column always stores ID as value
	idx := &entity.IndexInt{Slot: c.Slot}
	v, err := c.EvaluateCell(val, nil)
	if err != nil {
		return nil, err
	}

	idx.Val = v.(int64)
	return idx, nil
}

func (sequence) IndexKind() entity.Index {
	return &entity.IndexInt{}
}

func (sequence) GenerateValue() (string, error) {
	return strconv.FormatInt(snowflakeID(), 10), nil
}
