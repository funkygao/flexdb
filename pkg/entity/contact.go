package entity

// Contact is member of a team.
type Contact struct {
	TeamID int64 `gorm:"column:team_id" json:"team_id,omitempty"`

	ID     int64  `gorm:"column:id; AUTO_INCREMENT" json:"id,omitempty"`
	Name   string `gorm:"column:name" json:"name,omitempty"`
	Remark string `gorm:"column:remark" json:"remark,omitempty"`

	TeamRole TeamRole `gorm:"column:team_role" json:"team_role,omitempty"`
}
