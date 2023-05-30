package entity

import "time"

// Org is unit of tenant.
type Org struct {
	ID int64 `gorm:"column:id; AUTO_INCREMENT" json:"id"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	MTime time.Time `gorm:"column:mtime; type:timestamp; default: NOW();" json:"mtime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
	MUser string    `gorm:"column:muser" json:"muser,omitempty"`

	Name       string `gorm:"column:name" json:"name"`
	Secret     string `gorm:"column:secret" json:"-"`
	APIVersion string `gorm:"column:api_ver" json:"api_ver"`

	// quota related info
}
