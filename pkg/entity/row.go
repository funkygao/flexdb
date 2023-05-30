package entity

import (
	"reflect"
	"time"
)

// Row represent a row in underlying physical table.
// Row will not be directly exported, but converted into dto.RowData, which can be exported.
type Row struct {
	ModelID int64 `gorm:"column:model_id"` // Field(0)

	S1  string `gorm:"column:s1"` // Field(1) S1 MUST assure it is 1st field of struct Row: reflect.Field(1)
	S2  string `gorm:"column:s2"`
	S3  string `gorm:"column:s3"`
	S4  string `gorm:"column:s4"`
	S5  string `gorm:"column:s5"`
	S6  string `gorm:"column:s6"`
	S7  string `gorm:"column:s7"`
	S8  string `gorm:"column:s8"`
	S9  string `gorm:"column:s9"`
	S10 string `gorm:"column:s10"`
	S11 string `gorm:"column:s11"`
	S12 string `gorm:"column:s12"`
	S13 string `gorm:"column:s13"`
	S14 string `gorm:"column:s14"`
	S15 string `gorm:"column:s15"`
	S16 string `gorm:"column:s16"`
	S17 string `gorm:"column:s17"`
	S18 string `gorm:"column:s18"`
	S19 string `gorm:"column:s19"`
	S20 string `gorm:"column:s20"`
	S21 string `gorm:"column:s21"`
	S22 string `gorm:"column:s22"`
	S23 string `gorm:"column:s23"`
	S24 string `gorm:"column:s24"`
	S25 string `gorm:"column:s25"`
	S26 string `gorm:"column:s26"`
	S27 string `gorm:"column:s27"`
	S28 string `gorm:"column:s28"`
	S29 string `gorm:"column:s29"`
	S30 string `gorm:"column:s30"`
	S31 string `gorm:"column:s31"`
	S32 string `gorm:"column:s32"`
	S33 string `gorm:"column:s33"`
	S34 string `gorm:"column:s34"`
	S35 string `gorm:"column:s35"`
	S36 string `gorm:"column:s36"`
	S37 string `gorm:"column:s37"`
	S38 string `gorm:"column:s38"`
	S39 string `gorm:"column:s39"`
	S40 string `gorm:"column:s40"`
	S41 string `gorm:"column:s41"`
	S42 string `gorm:"column:s42"`
	S43 string `gorm:"column:s43"`
	S44 string `gorm:"column:s44"`
	S45 string `gorm:"column:s45"`
	S46 string `gorm:"column:s46"`
	S47 string `gorm:"column:s47"`
	S48 string `gorm:"column:s48"`
	S49 string `gorm:"column:s49"`
	S50 string `gorm:"column:s50"`

	R1  string `gorm:"column:r1"`  // reserved by column plugin
	R2  string `gorm:"column:r2"`  // reserved by column plugin
	R3  string `gorm:"column:r3"`  // reserved by column plugin
	R4  string `gorm:"column:r4"`  // reserved by column plugin
	R5  string `gorm:"column:r5"`  // reserved by column plugin
	R6  string `gorm:"column:r6"`  // reserved by column plugin
	R7  string `gorm:"column:r7"`  // reserved by column plugin
	R8  string `gorm:"column:r8"`  // reserved by column plugin
	R9  string `gorm:"column:r9"`  // reserved by column plugin
	R10 string `gorm:"column:r10"` // reserved by column plugin

	// DO NOT move it elsewhere!
	// Otherwise, unit test will break.
	//
	// Yes, I know it should be placed above and that will look nicer...
	// BUT for performance: reflect can reference field by id instead of name, ID field must be here!
	ID uint64 `gorm:"primaryKey;autoIncrement:true"`

	Memo string `gorm:"column:memo"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	MTime time.Time `gorm:"column:mtime; type:timestamp; default: NOW();" json:"mtime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
	MUser string    `gorm:"column:muser" json:"muser,omitempty"`

	Deleted bool `gorm:"column:deleted"`

	// Ver is the optimistic lock
	Ver int `gorm:"column:ver" json:"ver"`

	//=============
	// associations
	//=============

	Clob *RowClob `gorm:"foreignKey:RowID"`

	Changelogs []RowChangelog `gorm:"foreignKey:RowID"`

	//==========
	// transient
	//==========

	orgID int64 `gorm:"-"`
}

// SetField set a string value on specified slot.
func (r *Row) SetField(slot int16, val string) {
	if slot > maxSlots || slot < firstSlot {
		return
	}

	v := reflect.ValueOf(r).Elem().Field(int(slot))
	if v.IsValid() {
		v.SetString(val)
	}
}

// GetField returns string value of the specified slot.
func (r *Row) GetField(slot int16) string {
	if slot > maxSlots || slot < firstSlot {
		return ""
	}

	v := reflect.ValueOf(r).Elem().Field(int(slot))
	if !v.IsValid() {
		return ""
	}

	return v.String()
}

// HasClob tells whether the row has storage in RowClob.
func (r *Row) HasClob() bool {
	return r.Clob != nil && r.Clob.RowID == r.ID
}

// evaluateCell extract value from row/clob or calculator according to column.
func (r *Row) evaluateCell(c Column) interface{} {
	if c.ClobWise() {
		// column might be stored in clob
		if r.Clob == nil {
			// did not insert row for this column
			return ""
		}

		// clob can never be value calculator
		return r.Clob.GetField(c.clobSlot())
	}

	val, _ := c.Plugin().EvaluateCell(r.GetField(c.Slot), r)
	return val
}
