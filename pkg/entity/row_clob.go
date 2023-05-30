package entity

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

// RowClob character large objects is an extension of Row.
// When retrieving data, Row will join Clob IfNec.
type RowClob struct {
	RowID uint64 `gorm:"column:row_id"`

	S1 string `gorm:"column:s1"` // Field(1) S1 MUST assure it is 1st field of struct RowClob: reflect.Field(1)
	S2 string `gorm:"column:s2"` // Field(2)
	S3 string `gorm:"column:s3"` // Field(3)

	Deleted bool `gorm:"column:deleted" json:"-"` // public invisible

	//==========
	// transient
	//==========

	orgID int64 `gorm:"-"`
}

// TableName is a gorm hook.
func (c RowClob) TableName() string {
	//return "Clob" + strconv.FormatInt(c.orgID, 10)
	return "Clob"
}

// BeforeCreate is a gorm hook.
func (c *RowClob) BeforeCreate(tx *gorm.DB) error {
	if c.orgID < 1 {
		return ErrClobEmptyOrgID
	}
	if c.RowID < 1 {
		return ErrClobEmptyRowID
	}

	for slot := firstSlot; slot <= maxClobSlots; slot++ {
		if len(c.GetField(slot)) > clobColumnMaxSize {
			return fmt.Errorf("Clob slot[%d] max length reached", slot)
		}
	}

	return nil
}

func (c *RowClob) IsZero() bool {
	return c.S1 == "" && c.S2 == "" && c.S3 == ""
}

func (c *RowClob) SetField(slot int16, val string) {
	if slot > maxClobSlots || slot < firstSlot {
		return
	}

	v := reflect.ValueOf(c).Elem().Field(int(slot))
	if v.IsValid() {
		v.SetString(val)
	}
}

func (c *RowClob) GetField(slot int16) string {
	if slot > maxClobSlots || slot < firstSlot {
		return ""
	}

	v := reflect.ValueOf(c).Elem().Field(int(slot))
	if !v.IsValid() {
		return ""
	}

	return v.String()
}
