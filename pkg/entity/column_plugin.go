package entity

// ColumnPlugin is the required plugin spec of the meta column framework.
type ColumnPlugin interface {

	// ValidateCell assures that cell value is valid before insert/update data.
	ValidateCell(val string) error

	// EvaluateCell evaluates string format cell value into corresponding golang data type.
	// e,g. ColumnDatetime evaluates into time.Time
	EvaluateCell(cell string, row *Row) (interface{}, error)
}

// Introspector is an optional column plugin that validates the column itself when add/update the column.
// In most cases, Column validates itself enough, but if you want more, implement Introspector.
type Introspector interface {

	// Introspect validate the column itself when add/update the column.
	Introspect() error
}

// Indexable is an optional column plugin that helps create index for data.
type Indexable interface {
	// CreateIndex creates a concrete index when inserting row data.
	CreateIndex(cell string) (Index, error)

	// IndexKind returns the actual index type.
	IndexKind() Index
}

// CellValueGenerator is an optional column plugin that can auto generate cell value and persists in row.
type CellValueGenerator interface {
	GenerateValue() (string, error)
}
