package entity

import "time"

// Share stores an app can be accessed by whom.
type Share struct {
	// AppID specifies which app this model belongs to.
	AppID int64 `gorm:"column:app_id" json:"app_id"`

	ID      int64  `gorm:"column:id; AUTO_INCREMENT" json:"id,omitempty"`
	Subject string `gorm:"column:subject" json:"subject,omitempty"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
}
