package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/lithammer/shortuuid/v3"
	"gorm.io/gorm"
)

// TableName is a gorm hook.
func (App) TableName() string {
	return "App"
}

// BeforeCreate is a gorm hook.
func (a *App) BeforeCreate(tx *gorm.DB) (err error) {
	if a.OrgID < 1 {
		return ErrAppEmptyOrgID
	}
	if err = a.validateName(); err != nil {
		return err
	}

	a.Status = AppStatusInit
	now := time.Now()
	if a.CTime.IsZero() {
		a.CTime = now
	}
	if a.MTime.IsZero() {
		a.MTime = now
	}
	if a.UUID == "" {
		a.UUID = shortuuid.New()
	}
	return
}

// BeforeUpdate is a gorm hook.
func (a *App) BeforeUpdate(tx *gorm.DB) (err error) {
	if a.ID == 0 {
		return ErrAppEmptyID
	}
	if _, present := appStatusLabels[a.Status]; !present {
		return fmt.Errorf("invalid app.status:%v", a.Status)
	}

	return
}

func (a *App) validateName() error {
	if strings.TrimSpace(a.Name) == "" {
		return ErrAppEmptyName
	}
	return nil
}
