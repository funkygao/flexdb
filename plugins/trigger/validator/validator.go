// Package validator provides model level validation based on rule.
package validator

import (
	"github.com/agile-app/flexdb/pkg/entity"
)

type validator struct {
}

func (validator) Name() string {
	return "validator"
}

func (validator) ID() int {
	return 4
}

func (t *validator) BeforeInsert(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *validator) AfterInsert(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *validator) BeforeUpdate(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *validator) AfterUpdate(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *validator) BeforeDelete(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}

func (t *validator) AfterDelete(c entity.TriggerContext, m *entity.Model, r *entity.Row, action string) error {
	return nil
}
