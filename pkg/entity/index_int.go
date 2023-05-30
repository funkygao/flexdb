package entity

import (
	"gorm.io/gorm"
)

// IndexInt is index for int data.
type IndexInt struct {
	ModelID int64 `gorm:"column:model_id"`

	ID   uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Slot int16  `gorm:"column:slot"`
	Val  int64  `gorm:"column:val"`

	RowID uint64 `gorm:"column:row_id"`

	//==========
	// transient
	//==========

	unique bool `gorm:"-"`
}

// TableName is a gorm hook.
func (i IndexInt) TableName() string {
	//return "IndexInt_" + strconv.FormatInt(i.orgID, 10)
	if i.unique {
		return "IndexIntUniq"
	}

	return "IndexInt"
}

// BeforeCreate is a gorm hook.
func (i IndexInt) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ModelID < 1 {
		return ErrIndexEmptyModelID
	}

	return
}

func (i *IndexInt) Unique() bool {
	return i.unique
}

func (i *IndexInt) SlotID() int16 {
	return i.Slot
}

func (i *IndexInt) SetRowID(id uint64) {
	i.RowID = id
}

func (i *IndexInt) GetRowID() uint64 {
	return i.RowID
}

func (i *IndexInt) setModelID(modelID int64) {
	i.ModelID = modelID
}

func (i *IndexInt) setUnique(unique bool) {
	i.unique = unique
}

func (i *IndexInt) Value() interface{} {
	return i.Val
}
