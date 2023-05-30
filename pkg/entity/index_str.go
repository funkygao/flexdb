package entity

import (
	"gorm.io/gorm"
)

// IndexStr is index for str data.
type IndexStr struct {
	ModelID int64 `gorm:"column:model_id"`

	ID   uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Slot int16  `gorm:"column:slot"`
	Val  string `gorm:"column:val"`

	RowID uint64 `gorm:"column:row_id"`

	//==========
	// transient
	//==========

	unique bool `gorm:"-"`
}

func (i *IndexStr) Unique() bool {
	return i.unique
}

// TableName is a gorm hook.
func (i IndexStr) TableName() string {
	//return "IndexStr_" + strconv.FormatInt(i.orgID, 10)
	if i.unique {
		return "IndexStrUniq"
	}

	return "IndexStr"
}

// BeforeCreate is a gorm hook.
func (i IndexStr) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ModelID < 1 {
		return ErrIndexEmptyModelID
	}

	return
}

func (i *IndexStr) SetRowID(id uint64) {
	i.RowID = id
}

func (i *IndexStr) SlotID() int16 {
	return i.Slot
}

func (i *IndexStr) GetRowID() uint64 {
	return i.RowID
}

func (i *IndexStr) setModelID(modelID int64) {
	i.ModelID = modelID
}

func (i *IndexStr) setUnique(unique bool) {
	i.unique = unique
}

func (i *IndexStr) Value() interface{} {
	return i.Val
}
