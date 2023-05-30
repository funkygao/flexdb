package entity

import "errors"

// Team manages users of an app.
type Team struct {
	// AppID specifies which app this model belongs to.
	AppID int64 `gorm:"column:app_id" json:"app_id,omitempty"`

	ID     int64  `gorm:"column:id; AUTO_INCREMENT" json:"id,omitempty"`
	Name   string `gorm:"column:name" json:"name,omitempty"`
	Remark string `gorm:"column:remark" json:"remark,omitempty"`

	//=============
	// associations
	//=============

	Contacts []Contact `gorm:"foreignKey:TeamID" json:"-"`
}

func (t *Team) AddMember(c Contact) error {
	if len(t.Contacts) > maxMembersPerTeam {
		return errors.New("too many members")
	}

	t.Contacts = append(t.Contacts, c)
	return nil
}
