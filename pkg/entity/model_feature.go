package entity

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
)

// IMPORTANT: featureSize cannot be changed once set!
const featureSize = 100

var _ sql.Scanner = (*ModelFeature)(nil)

// ModelFeature is features of a model.
type ModelFeature [featureSize]byte

// Label renders model feature in human readable content.
func (f *ModelFeature) Label() string {
	s := []string{}
	if f.ChangeAuditEnabled() {
		s = append(s, "A")
	}
	if f.FakeDeleteEnabled() {
		s = append(s, "F")
	}
	if f.CreateRowEnabled() {
		s = append(s, "C")
	}
	if f.ReadRowEnabled() {
		s = append(s, "R")
	}
	if f.UpdateRowEnabled() {
		s = append(s, "U")
	}
	if f.DeleteRowEnabled() {
		s = append(s, "D")
	}
	if len(s) == 0 {
		return "-"
	}

	return strings.Join(s, "")
}

// Scan implements sql.Scanner so Feature can be read from databases transparently.
func (f *ModelFeature) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		// if an empty Feature comes from a table, we return a null Feature
		if src == "" {
			return nil
		}

		f1, err := f.parseFeature(src)
		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}

		*f = f1

	case []byte:
		// if an empty UUID comes from a table, we return a null UUID
		if len(src) == 0 {
			return nil
		}

		// assumes a simple slice of bytes if {featureSize} bytes
		// otherwise attempts to parse
		if len(src) != featureSize {
			return f.Scan(string(src))
		}

		copy((*f)[:], src)

	default:
		return fmt.Errorf("Scan: unable to scan type %T into Feature", src)
	}

	return nil
}

func (ModelFeature) parseFeature(s string) (ModelFeature, error) {
	var f ModelFeature
	for i, b := range s {
		f[i] = byte(b)
	}
	return f, nil
}

// Value implements sql.Valuer so that Feature can be written to databases
// transparently. Currently, ModelFeature map to string.
func (f ModelFeature) Value() (driver.Value, error) {
	return string(f[:]), nil
}
