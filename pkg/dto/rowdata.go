package dto

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/agile-app/flexdb/internal/spec"
)

// RowData is the DTO of row.
type RowData map[string]interface{}

// ID returns id of the row data.
func (r RowData) ID() uint64 {
	id, _ := strconv.ParseUint(r.StrValueOf(KeyID), 10, 64)
	return id
}

// SetID set id of the row data.
func (r RowData) SetID(id uint64) {
	r[KeyID] = strconv.FormatUint(id, 10)
}

// CTime returns ctime value.
func (r RowData) CTime() time.Time {
	t, _ := time.Parse(spec.SimpleDateFormat, r.StrValueOf(KeyCTime))
	return t
}

// SetCTime set ctime value.
func (r RowData) SetCTime(t time.Time) {
	r[KeyCTime] = t.Format(spec.SimpleDateFormat)
}

// MTime returns mtime value.
func (r RowData) MTime() time.Time {
	t, _ := time.Parse(spec.SimpleDateFormat, r.StrValueOf(KeyMTime))
	return t
}

// SetMTime set mtime value.
func (r RowData) SetMTime(t time.Time) {
	r[KeyMTime] = t.Format(spec.SimpleDateFormat)
}

// CUser returns cuser value.
func (r RowData) CUser() string {
	return r.StrValueOf(KeyCUser)
}

// SetCUser set cuser value.
func (r RowData) SetCUser(u string) {
	r[KeyCUser] = u
}

// MUser returns muser value.
func (r RowData) MUser() string {
	return r.StrValueOf(KeyMUser)
}

// SetMUser set muser value.
func (r RowData) SetMUser(u string) {
	r[KeyMUser] = u
}

// Memo returns memo value.
func (r RowData) Memo() string {
	return r.StrValueOf(KeyMemo)
}

// SetMemo set memo value.
func (r RowData) SetMemo(v string) {
	r[KeyMemo] = v
}

// Put set value by key.
func (r RowData) Put(k string, v interface{}) {
	r[k] = v
}

// StrValueOf gets value by key.
// As we store cell data as varchar, data will be converted into string before persistence.
func (r RowData) StrValueOf(k string) string {
	switch v := r[k].(type) {
	case string:
		if v == zeroTimeStr {
			// empty time
			return ""
		}
		return strings.TrimSpace(v)
	case float64:
		return strconv.FormatInt(int64(v), 10)
	case nil:
		return ""
	case bool:
		return strconv.FormatBool(v)
	case int64:
		return strconv.FormatInt(v, 10)
	default:
		panic(fmt.Sprintf("Unknown type: %T", v))
	}
}

// HasField tells whether a key exists.
func (r RowData) HasField(k string) bool {
	if _, present := r[k]; present {
		return true
	}

	return false
}

// JSONString renders row data in json.
func (r RowData) JSONString() string {
	b, _ := json.Marshal(r)
	return string(b)
}
