package entity

type Task struct {
	AppID int64 `gorm:"column:app_id" json:"-"`

	ID int64 `gorm:"column:id; AUTO_INCREMENT" json:"id,omitempty"`

	Status int16  `gorm:"column:status" json:"status,omitempty"`
	Body   string `gorm:"column:body" json:"body,omitempty"`
}

// TableName is a gorm hook.
func (t Task) TableName() string {
	return "Task"
}
