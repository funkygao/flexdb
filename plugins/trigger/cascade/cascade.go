package cascade

import (
	"strconv"
	"strings"

	"github.com/agile-app/flexdb/pkg/entity"
)

type cascade struct {
}

func (cascade) Name() string {
	return "cascade"
}

func (cascade) ID() int {
	return 2
}

func (t *cascade) BeforeInsert(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *cascade) AfterInsert(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	tuple := strings.SplitN(action, ",", 4)
	sourceColName, targetModel, targetRow, targetCol := tuple[0], tuple[1], tuple[2], tuple[3]
	targetModelID, err := strconv.ParseInt(targetModel, 10, 64)
	if err != nil {
		return err
	}

	tm, err := c.LoadModel(targetModelID)
	if err != nil {
		return err
	}

	targetColOk := false
	for _, c := range tm.EntityModel().Slots {
		if c.Name == targetCol {
			// validate ok
			targetColOk = true
			break
		}
	}
	if !targetColOk {
		// TODO trigger error. rollback?
	}

	tid, err := strconv.ParseUint(targetRow, 10, 64)
	if err != nil {
		return err
	}

	targetRowData, err := tm.RetrieveRow(tid, false)
	if err != nil {
		return err
	}

	sourceCol := m.SlotByName(sourceColName)
	targetRowData.Master.Put(targetCol, r.GetField(sourceCol.Slot))
	return nil
}

func (t *cascade) BeforeUpdate(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *cascade) AfterUpdate(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *cascade) BeforeDelete(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *cascade) AfterDelete(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}
