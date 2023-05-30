package entity

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// TableName is a gorm hook.
func (Column) TableName() string {
	return "Field"
}

// BeforeCreate is a gorm hook.
func (c *Column) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ModelID < 1 {
		return ErrColumnEmptyModelID
	}
	if c.Kind == unknownColumnKind {
		return ErrColumnUnknownKind
	}

	if c.Slot < 1 {
		return ErrColumnEmptySlot
	}
	if c.Slot > maxSlots && !c.ClobWise() {
		return fmt.Errorf("slot[%d] too big", c.Slot)
	}
	if c.ClobWise() {
		if c.clobSlot() > maxClobSlots {
			return fmt.Errorf("clob slot too big")
		}
	}
	if strings.TrimSpace(c.Name) == "" {
		return ErrColumnEmptyName
	}
	if err = c.validateName(); err != nil {
		return err
	}
	if c.Ordinal < 1 {
		c.Ordinal = c.Slot
	}
	if c.Kind != ColumnChoice {
		c.Choices = ""
	} else {
		// strip empty comma
		c.Choices = strings.Trim(c.Choices, ",")
	}

	now := time.Now()
	if c.CTime.IsZero() {
		c.CTime = now
	}
	if c.MTime.IsZero() {
		c.MTime = now
	}

	return
}

// BeforeUpdate is hook of gorm.
func (c *Column) BeforeUpdate(tx *gorm.DB) (err error) {
	if c.ID < 1 {
		return ErrColumnEmptyID
	}
	if c.ModelID < 1 {
		return ErrColumnEmptyModelID
	}

	if c.Name != "" {
		if err = c.validateName(); err != nil {
			return err
		}
	}
	if c.Kind != ColumnChoice {
		c.Choices = ""
	} else {
		// strip empty comma
		c.Choices = strings.Trim(c.Choices, ",")
	}

	if c.MTime.IsZero() {
		c.MTime = time.Now()
	}

	return
}

// BeforeDelete is hook of gorm.
func (c *Column) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind is hook of gorm.
func (c *Column) AfterFind(tx *gorm.DB) (err error) {
	// TODO unmarshal widget related info
	return
}

func (c *Column) validateName() error {
	if strings.HasSuffix(c.Name, ReservedColumnNameSuffix) {
		return fmt.Errorf("%s cannot suffix with %s", c.Name, ReservedColumnNameSuffix)
	}

	for _, n := range illegalNames {
		if strings.Contains(c.Name, n) {
			return fmt.Errorf("%s cannot contain %s", c.Name, n)
		}
	}

	return nil
}
