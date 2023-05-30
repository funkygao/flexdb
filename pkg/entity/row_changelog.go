package entity

import "time"

// RowChangelog records changelog of a row: audit.
type RowChangelog struct {
	RowID uint64 `gorm:"column:row_id"`

	Diff string `gorm:"column:diff" json:"diff"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
}

// TableName is a gorm hook.
func (RowChangelog) TableName() string {
	return "DataChangelog"
}
