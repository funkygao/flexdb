package many2one

import (
	"fmt"
	"strconv"

	"github.com/agile-app/flexdb/pkg/entity"
)

const (
	reservedColumn int = 2
)

// plugin compliance check
var (
	_ entity.ColumnPlugin       = (*many2one)(nil)
	_ entity.ReferenceValidator = (*many2one)(nil)
	_ entity.Indexable          = (*many2one)(nil)
)

type many2one struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *many2one) ValidateReference(m *entity.Model, toAddColumns []*entity.Column) error {
	if c.RefModelID == 0 {
		return fmt.Errorf("column[%s] ref model is null", c.Name)
	}

	if _, err := c.ctx.ModelAccessorOf(c.RefModelID); err != nil { // TODO perf shallow
		return err
	}

	// more than 1 many2one column ref the same model?
	for _, slot := range toAddColumns {
		if slot.Name == c.Name {
			// ignore self
			continue
		}

		// 1 model cannot have more than one many2one columns referencing to the same model
		if slot.Kind == c.Kind && slot.RefModelID == c.RefModelID {
			return fmt.Errorf("column: %s and %s ref same model", slot.Name, c.Name)
		}
	}

	return nil
}

func (c *many2one) ValidateCell(refRowIDstr string) error {
	refRowID, err := c.EvaluateCell(refRowIDstr, nil)
	if err != nil {
		return err
	}

	if id := refRowID.(uint64); id < 1 {
		return fmt.Errorf("empty value on column:%s", c.Name)
	}

	// TODO lookup ref table row to see whether it exists

	return nil
}

func (many2one) EvaluateCell(refRowIDstr string, row *entity.Row) (interface{}, error) {
	return strconv.ParseUint(refRowIDstr, 10, 64)
}

func (c *many2one) CreateIndex(refRowIDstr string) (entity.Index, error) {
	// always index this column
	idx := &entity.IndexInt{Slot: c.Slot}
	v, err := c.EvaluateCell(refRowIDstr, nil)
	if err != nil {
		return nil, err
	}

	idx.Val = int64(v.(uint64))
	return idx, nil
}

func (many2one) IndexKind() entity.Index {
	return &entity.IndexInt{}
}

func (c *many2one) CellValue(r *entity.Row) (interface{}, error) {
	refRowID, err := c.EvaluateCell(r.GetField(c.Slot), nil)
	if err != nil {
		return "", err
	}

	targetModel, err := c.PluginContext.ModelAccessorOf(c.RefModelID)
	if err != nil {
		return "", err
	}

	md, err := targetModel.RetrieveRow(refRowID.(uint64), false)
	if err != nil {
		return "", err
	}

	return md.Master.StrValueOf(targetModel.EntityModel().IdentName), nil
}
