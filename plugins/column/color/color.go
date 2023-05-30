package color

import (
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin            = (*color)(nil)
	_ entity.Indexable               = (*color)(nil)
	_ entity.ListViewWidgetTypeAware = (*color)(nil)
)

type color struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *color) ValidateCell(val string) error {
	return nil
}

func (color) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *color) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (color) IndexKind() entity.Index {
	return &entity.IndexStr{}
}

func (color) ListViewWidgetType() string {
	return amis.ATColor
}
