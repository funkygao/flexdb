package dto

import (
	"strconv"
	"strings"
)

// QuickSave is the POST body of a quick save form.
type QuickSave struct {
	Rows     []RowData           `json:"rows"`
	RowsDiff []map[string]string `json:"rowsDiff"`
	IDs      string              `json:"ids"` // comma separated id, e,g. 21,56,99
}

func (qs QuickSave) RowIDs() (ids []uint64) {
	tuple := strings.Split(qs.IDs, ",")
	ids = make([]uint64, 0, len(tuple))
	for _, item := range tuple {
		if id, err := strconv.ParseUint(item, 10, 64); err == nil {
			ids = append(ids, id)
		}
	}

	return
}
