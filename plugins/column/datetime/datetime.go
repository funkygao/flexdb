package datetime

import (
	"strconv"
	"time"

	"github.com/agile-app/flexdb/internal/spec"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin            = (*datetime)(nil)
	_ entity.Indexable               = (*datetime)(nil)
	_ entity.ListViewExtraCeller     = (*datetime)(nil)
	_ entity.ListViewWidgetTypeAware = (*datetime)(nil)
)

type datetime struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *datetime) ValidateCell(val string) error {
	_, err := c.EvaluateCell(val, nil)
	return err
}

func (datetime) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	// amis passes timestamp
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(i, 0), nil
}

func (c *datetime) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed || val == "" {
		return nil, nil
	}

	idx := &entity.IndexTime{Slot: c.Slot}

	t, err := time.Parse(spec.SimpleDateFormat, val)
	if err != nil {
		return nil, err
	}

	idx.Val = t
	return idx, nil
}

func (datetime) IndexKind() entity.Index {
	return &entity.IndexTime{}
}

func (datetime) ListViewWidgetType() string {
	return amis.ATDate
}

func (c datetime) ListViewExtraCell(currentCellVal string) (extraCellColumn, extraCellVal string) {
	extraCellColumn = c.Name + entity.ReservedColumnNameSuffix
	if currentCellVal == "" {
		return extraCellColumn, ""
	}

	t, err := c.EvaluateCell(currentCellVal, nil)
	if err != nil {
		return extraCellColumn, ""
	}

	return extraCellColumn, t.(time.Time).Format(spec.YYYYMMDD)
}
