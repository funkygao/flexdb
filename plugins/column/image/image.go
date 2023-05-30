package image

import (
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin            = (*image)(nil)
	_ entity.ListViewWidgetTypeAware = (*image)(nil)
)

type image struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *image) ValidateCell(val string) error {
	return nil
}

func (image) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (image) ListViewWidgetType() string {
	return amis.ATImage
}
