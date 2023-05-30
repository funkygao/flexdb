package entity

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// PickItem works with ColumnChoice and stores its enum values.
type PickItem struct {
	ModelID int64 `gorm:"column:model_id" json:"-"`

	ID   int64  `gorm:"column:id; AUTO_INCREMENT" json:"val"`
	Slot int16  `gorm:"column:slot" json:"-"`
	Val  string `gorm:"column:val" json:"label"`

	// ParentID used in chained select scenarios.
	ParentID sql.NullInt64 `gorm:"column:parent_id" json:"-"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	MTime time.Time `gorm:"column:mtime; type:timestamp; default: NOW();" json:"mtime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
	MUser string    `gorm:"column:muser" json:"muser,omitempty"`
}

// TableName is a gorm hook.
func (PickItem) TableName() string {
	return "Picklist"
}

// BeforeCreate is hook of gorm.
func (p *PickItem) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ModelID < 1 {
		return errors.New("empty pickitem.modelID")
	}
	if p.Slot < 1 {
		return errors.New("empty pickitem.slot")
	}
	if strings.TrimSpace(p.Val) == "" {
		return errors.New("empty pickitem.value")
	}
	return
}

// BeforeUpdate is hook of gorm.
func (p *PickItem) BeforeUpdate(tx *gorm.DB) (err error) {
	if p.ID < 1 {
		return errors.New("empty pickitem.id")
	}
	// TODO only Val can change
	return
}
