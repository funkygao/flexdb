package entity

// Relation is the relationship between models.
type Relation struct {
	ID int64 `gorm:"column:id; AUTO_INCREMENT" json:"id"`

	RefKind ColumnKind `gorm:"column:ref_kind" json:"ref_kind"`

	FromModelID int64 `gorm:"column:from_model_id" json:"from_model_id,omitempty"`
	FromSlot    int16 `gorm:"column:from_slot" json:"from_slot,omitempty"`

	ToModelID int64 `gorm:"column:to_model_id" json:"to_model_id,omitempty"`
	ToSlot    int16 `gorm:"column:to_slot" json:"to_slot,omitempty"`
}

// TableName is a gorm hook.
func (Relation) TableName() string {
	return "Relationship"
}
