package entity

import "time"

// A Chart belongs to an App.
type Chart struct {
	AppID int64 `gorm:"column:app_id" json:"app_id,omitempty"`

	ID   int64  `gorm:"column:id; AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name" json:"name,omitempty"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	MTime time.Time `gorm:"column:mtime; type:timestamp; default: NOW();" json:"mtime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
	MUser string    `gorm:"column:muser" json:"muser,omitempty"`
}
