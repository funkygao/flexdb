package entity

var (
	_ Index = (*IndexInt)(nil)
	_ Index = (*IndexStr)(nil)
	_ Index = (*IndexTime)(nil)
)

// Index is abstraction of secondary index with all types of concrete indexes.
type Index interface {
	// TableName returns underlying table name of the index.
	TableName() string

	Unique() bool

	SetRowID(id uint64)
	GetRowID() uint64

	SlotID() int16

	// Value returns value of the index row.
	// According data type of the index, result can be string/int/time/etc.
	Value() interface{}

	setUnique(bool)
	setModelID(int64)
}
