package workflow

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

type workflow struct {
}

func (workflow) Name() string {
	return "workflow"
}

func (workflow) ID() int {
	return 3
}

func (t *workflow) BeforeInsert(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *workflow) AfterInsert(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *workflow) BeforeUpdate(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *workflow) AfterUpdate(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *workflow) BeforeDelete(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *workflow) AfterDelete(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}
