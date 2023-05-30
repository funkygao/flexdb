package entity

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Page belongs to an app that renders UI.
type Page struct {
	AppID int64 `gorm:"column:app_id" json:"-"`

	ID   int64  `gorm:"AUTO_INCREMENT" json:"id"`
	Name string `json:"name"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	MTime time.Time `gorm:"column:mtime; type:timestamp; default: NOW();" json:"mtime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
	MUser string    `gorm:"column:muser" json:"muser,omitempty"`
}

// TableName is a gorm hook.
func (Page) TableName() string {
	return "AppPage"
}

// BeforeCreate is a gorm hook.
func (p *Page) BeforeCreate(tx *gorm.DB) (err error) {
	if p.AppID < 1 {
		return errors.New("empty page.app")
	}
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("empty page.name")
	}

	return
}
