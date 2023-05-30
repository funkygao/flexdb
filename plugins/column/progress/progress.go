package progress

import (
	"fmt"
	"strconv"

	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin            = (*progress)(nil)
	_ entity.ListViewWidgetTypeAware = (*progress)(nil)
)

type progress struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *progress) ValidateCell(val string) error {
	if v, err := c.EvaluateCell(val, nil); err != nil || v.(int) > 100 {
		return fmt.Errorf("%s: illegal progress value", val)
	}

	return nil
}

func (progress) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return strconv.Atoi(val)
}

func (progress) ListViewWidgetType() string {
	return amis.ATProgress
}
