package boolean

import (
	"strconv"

	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin            = (*boolean)(nil)
	_ entity.Introspector            = (*boolean)(nil)
	_ entity.ListViewWidgetTypeAware = (*boolean)(nil)

	mappings = map[string]bool{
		"是": true,
		"否": false,
		"非": false,
		"对": true,
		"错": false,
		"真": true,
		"假": false,
	}
)

type boolean struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *boolean) ValidateCell(val string) error {
	_, err := c.EvaluateCell(val, nil)
	return err
}

func (boolean) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	if v, present := mappings[val]; present {
		return v, nil
	}
	return strconv.ParseBool(val)
}

func (boolean) ListViewWidgetType() string {
	return amis.ATStatus
}

func (c *boolean) Introspect() error {
	c.Required = false // FIXME 目前amis switch控件的问题
	return nil
}
