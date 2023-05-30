package entity

import "fmt"

type triggerEvent int

const (
	TriggerBeforeCreate = triggerEvent(1)
	TriggerAfterCreate  = triggerEvent(2)
	TriggerBeforeUpdate = triggerEvent(3)
	TriggerAfterUpdate  = triggerEvent(4)
	TriggerBeforeDelete = triggerEvent(5)
	TriggerAfterDelete  = triggerEvent(6)
)

// ModelTrigger belongs to a model and acts as interceptor for persistence of data.
type ModelTrigger struct {
	ModelID int64 `gorm:"column:model_id" json:"-"`

	ID        int64 `gorm:"AUTO_INCREMENT" json:"id"`
	TriggerID int   `gorm:"column:trigger_id" json:"-"`

	BeforeInsertAction string `gorm:"column:before_insert" json:"before_insert"`
	AfterInsertAction  string `gorm:"column:after_insert" json:"after_insert"`
	BeforeUpdateAction string `gorm:"column:before_update" json:"before_update"`
	AfterUpdateAction  string `gorm:"column:after_update" json:"after_update"`
	BeforeDeleteAction string `gorm:"column:before_delete" json:"before_delete"`
	AfterDeleteAction  string `gorm:"column:after_delete" json:"after_delete"`
}

// TableName is a gorm hook.
func (ModelTrigger) TableName() string {
	return "ModelTrigger"
}

// InvokeBeforeInsert is a trigger hook.
func (t ModelTrigger) InvokeBeforeInsert(ctx TriggerContext, m *Model, r *Row) error {
	spec, present := triggers[int(t.ID)]
	if !present {
		return fmt.Errorf("Trigger:%d unknown", t.ID)
	}

	return spec.BeforeInsert(ctx, m, r, t.BeforeInsertAction)
}

// InvokeAfterInsert is a trigger hook.
func (t ModelTrigger) InvokeAfterInsert(ctx TriggerContext, m *Model, r *Row) error {
	spec, present := triggers[int(t.ID)]
	if !present {
		return fmt.Errorf("Trigger:%d unknown", t.ID)
	}

	return spec.AfterInsert(ctx, m, r, t.AfterInsertAction)
}

// InvokeBeforeUpdate is a trigger hook.
func (t ModelTrigger) InvokeBeforeUpdate(ctx TriggerContext, m *Model, r *Row) error {
	spec, present := triggers[int(t.ID)]
	if !present {
		return fmt.Errorf("Trigger:%d unknown", t.ID)
	}

	return spec.BeforeUpdate(ctx, m, r, t.BeforeUpdateAction)
}

// InvokeAfterUpdate is a trigger hook.
func (t ModelTrigger) InvokeAfterUpdate(ctx TriggerContext, m *Model, r *Row) error {
	spec, present := triggers[int(t.ID)]
	if !present {
		return fmt.Errorf("Trigger:%d unknown", t.ID)
	}

	return spec.AfterUpdate(ctx, m, r, t.AfterUpdateAction)
}

// InvokeBeforeDelete is a trigger hook.
func (t ModelTrigger) InvokeBeforeDelete(ctx TriggerContext, m *Model, r *Row) error {
	spec, present := triggers[int(t.ID)]
	if !present {
		return fmt.Errorf("Trigger:%d unknown", t.ID)
	}

	return spec.BeforeDelete(ctx, m, r, t.BeforeDeleteAction)
}

// InvokeAfterDelete is a trigger hook.
func (t ModelTrigger) InvokeAfterDelete(ctx TriggerContext, m *Model, r *Row) error {
	spec, present := triggers[int(t.ID)]
	if !present {
		return fmt.Errorf("Trigger:%d unknown", t.ID)
	}

	return spec.AfterDelete(ctx, m, r, t.AfterDeleteAction)
}
