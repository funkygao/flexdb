package entity

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// TableName is a gorm hook.
func (Model) TableName() string {
	return "Model"
}

// BeforeCreate is hook of gorm.
func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.AppID < 1 {
		return ErrModelEmptyAppID
	}
	if m.Kind == unknownModelKind {
		return ErrModelEmptyKind
	}
	if strings.TrimSpace(m.Name) == "" {
		return ErrModelEmptyName
	}
	if err = m.validateName(); err != nil {
		return err
	}

	m.Deleted = false
	m.Ver = 1
	now := time.Now()
	if m.CTime.IsZero() {
		m.CTime = now
	}
	if m.MTime.IsZero() {
		m.MTime = now
	}

	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	if m.AppID < 1 {
		return ErrModelEmptyAppID
	}
	if m.Name != "" {
		if err = m.validateName(); err != nil {
			return err
		}
	}

	return
}

func (m *Model) validateName() error {
	for _, n := range illegalNames {
		if strings.Contains(m.Name, n) {
			return fmt.Errorf("%s cannot contain %s", m.Name, n)
		}
	}

	return nil
}
