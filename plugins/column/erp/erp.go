package erp

import (
	"fmt"

	"github.com/agile-app/flexdb/pkg/entity"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin   = (*erp)(nil)
	_ entity.Indexable      = (*erp)(nil)
	_ entity.ListViewHrefer = (*erp)(nil)
)

type erp struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *erp) ValidateCell(val string) error {
	return nil
}

func (erp) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return val, nil
}

func (c *erp) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
}

func (erp) IndexKind() entity.Index {
	return &entity.IndexStr{}
}

func (c erp) ListViewHref() (href interface{}, anchor interface{}) {
	return fmt.Sprintf("timline://chat/?topin=${%s}", c.Name), fmt.Sprintf("${%s}", c.Name)
}
