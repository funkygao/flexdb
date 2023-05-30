package entity

// TableName is a gorm hook.
func (Share) TableName() string {
	return "Share"
}
