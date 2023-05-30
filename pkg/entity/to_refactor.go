package entity

// ReferenceValidator is an optional column plugin that validates the column itself when add the column.
// Used when we add/update a column to a model.
type ReferenceValidator interface {

	// ValidateReference validate the column itself.
	ValidateReference(m *Model, toAddColumns []*Column) error
}

// Builtin tells if the model is predefined.
func (m *Model) Builtin() bool {
	return m.Kind == ModelPerm
}

// ListViewable tells whether the column can be present in list view.
func (c *Column) ListViewable() bool {
	return !c.ClobWise()
}

// Virtual column will always leave slot of the row empty. Deprecated
func (c *Column) Virtual() bool {
	if _, yes := virtualColumns[c.Kind]; yes {
		return true
	}

	return false
}

// Referable tells whether the column can be referenced by columns of other model.
func (c *Column) Referable() bool {
	return !c.Virtual() && !c.Relational() && c.Indexable()
}

func (c *Column) listViewExtraCell(currentCellVal string) (extraCellColumn, extraCellVal string) {
	if p, yes := c.Plugin().(ListViewExtraCeller); yes {
		return p.ListViewExtraCell(currentCellVal)
	}

	return "", ""
}

// OnlyPinCanInput tells whether only current user can input this column data in create row form.
func (m *Model) OnlyPinCanInput(col *Column) bool {
	return m.SingleRowKind() && col.Kind == ColumnERP
}

// SingleRowKind tells whether the model can insert 1 row per user.
func (m *Model) SingleRowKind() bool {
	return m.Kind == ModelSingleRowPerUser
}

const (
	// ModelSingleRowPerUser is a model where creator can only insert 1 row and multiple updates.
	ModelSingleRowPerUser ModelKind = 2
)
