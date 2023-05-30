package entity

import (
	"time"

	"gorm.io/gorm"
)

// IndexTime is index for time data.
type IndexTime struct {
	ModelID int64 `gorm:"column:model_id"`

	ID   uint64    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Slot int16     `gorm:"column:slot"`
	Val  time.Time `gorm:"column:val"`

	RowID uint64 `gorm:"column:row_id"`

	//==========
	// transient
	//==========

	unique bool `gorm:"-"`
}

func (i *IndexTime) Unique() bool {
	return i.unique
}

// TableName is a gorm hook.
func (i IndexTime) TableName() string {
	//return "IndexTime_" + strconv.FormatInt(i.OrgID, 10)
	if i.unique {
		return "IndexTimeUniq"
	}

	return "IndexTime"
}

// BeforeCreate is a gorm hook.
func (i IndexTime) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ModelID < 1 {
		return ErrIndexEmptyModelID
	}

	return
}

func (i *IndexTime) SetRowID(id uint64) {
	i.RowID = id
}

func (i *IndexTime) GetRowID() uint64 {
	return i.RowID
}

func (i *IndexTime) SlotID() int16 {
	return i.Slot
}

func (i *IndexTime) setModelID(modelID int64) {
	i.ModelID = modelID
}

func (i *IndexTime) setUnique(unique bool) {
	i.unique = unique
}

func (i *IndexTime) Value() interface{} {
	return i.Val
}
