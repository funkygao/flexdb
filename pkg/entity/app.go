package entity

import (
	"time"
)

// App is an application belonging to an organization.
type App struct {
	OrgID int64 `json:"org_id"`

	ID          int64  `gorm:"column:id; AUTO_INCREMENT" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"desc"`
	Logo        string `gorm:"column:logo" json:"logo"`
	UUID        string `gorm:"column:uuid" json:"uuid"`

	Status     AppStatus  `gorm:"column:status" json:"status"`
	Visibility Visibility `gorm:"column:visibility" json:"visibility"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	MTime time.Time `gorm:"column:mtime; type:timestamp; default: NOW();" json:"mtime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
	MUser string    `gorm:"column:muser" json:"muser,omitempty"`

	//=============
	// associations
	//=============

	Models []Model `gorm:"foreignKey:AppID" json:"-"` // will not render in json
	Pages  []Page  `gorm:"foreignKey:AppID" json:"-"` // will not render in json
	Charts []Chart `gorm:"foreignKey:AppID" json:"-"` // will not render in json
	Shares []Share `gorm:"foreignKey:AppID" json:"-"` // will not render in json
}

// PermModels returns permission models of this app.
func (a *App) PermModels() []Model {
	r := make([]Model, 0, len(PermModels))
	for _, m := range a.Models {
		if m.Kind == ModelPerm {
			r = append(r, m)
		}
	}

	return r
}
