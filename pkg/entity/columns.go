package entity

import "sort"

// Columns is slice of Column used for sorting.
type Columns []*Column

var _ sort.Interface = (*Columns)(nil)

func (cols Columns) Len() int {
	return len(cols)
}

func (cols Columns) Less(i, j int) bool {
	return cols[i].Ordinal < cols[j].Ordinal
}

func (cols Columns) Swap(i, j int) {
	cols[i], cols[j] = cols[j], cols[i]
}
