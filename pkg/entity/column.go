package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/agile-app/flexdb/pkg/dto"
)

// Column is business attribute of a model.
// A model is made of builtin columns and user defined columns(what we call slot).
type Column struct {
	ModelID int64 `gorm:"column:model_id" json:"-"`

	ID      int64      `gorm:"column:id; AUTO_INCREMENT" json:"id,omitempty"`
	Name    string     `gorm:"column:name" json:"name"`
	Label   string     `gorm:"column:label" json:"label"`
	Remark  string     `gorm:"column:remark" json:"remark"`
	Kind    ColumnKind `gorm:"column:kind" json:"kind"`
	Slot    int16      `gorm:"column:slot" json:"-"`    // public invisible
	Ordinal int16      `gorm:"column:ordinal" json:"-"` // public invisible

	Indexed  bool `gorm:"column:indexed" json:"indexed"`
	Sortable bool `gorm:"column:sortable" json:"sortable"`
	Required bool `gorm:"column:required" json:"required"`
	Unique   bool `gorm:"column:uniq" json:"uniq"`
	ReadOnly bool `gorm:"column:ro" json:"ro"`

	// Default is the column default value if cell data is empty.
	Default string `gorm:"column:dftval" json:"dftval,omitempty"`

	// Choices stores comma separated choices for ColumnChoice.
	// Choices are not shared between columns: one column has its private options.
	Choices string `gorm:"column:choices" json:"choices"`

	// RefModelID hold the reference model if column type is: many2one/lookup.
	RefModelID int64 `gorm:"column:ref_model_id" json:"ref_model,omitempty"`

	// RefSlot used for many2one/lookup column TODO
	RefSlot int16 `gorm:"column:ref_slot" json:"ref_slot,omitempty"`

	// BizRule stores single column value business rule expression.
	BizRule string `gorm:"column:rule" json:"rule"`

	Deprecated bool `gorm:"column:deprecated" json:"-"` // public invisible

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	MTime time.Time `gorm:"column:mtime; type:timestamp; default: NOW();" json:"mtime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
	MUser string    `gorm:"column:muser" json:"muser,omitempty"`

	//==========
	// transient
	//==========

	// H is DTO storage for UI form that transports extra arbitrary data.
	// Name H is learned from web framework gin.H.
	H dto.H `gorm:"-" json:"h,omitempty"`

	// PluginContext provides context facility for column plugin.
	// Its up to upper layer to inject this value and used by column plugins.
	PluginContext PluginContext `gorm:"-" json:"-"`

	// cache of the plugin instance to avoid plugin factory create plugin instance more than once.
	pluginCache ColumnPlugin `gorm:"-"`
}

// KindLabel returns human readable column kind.
func (c *Column) KindLabel() string {
	return c.Kind.Label()
}

// ValidateCellData validates the cell value of the current column.
func (c *Column) ValidateCellData(val string) (err error) {
	// the shared validation rule
	if c.Required && c.Default == "" && strings.TrimSpace(val) == "" {
		return fmt.Errorf("column[%s] is required", c.Name)
	}

	if val == "" {
		return nil
	}

	// extension
	return c.Plugin().ValidateCell(val)
}

// ChoiceOptions returns enumerable options of the column from choices.
func (c *Column) ChoiceOptions() []string {
	if c.Kind != ColumnChoice {
		return nil
	}

	if c.Choices != "" {
		return strings.Split(strings.Trim(c.Choices, ","), ",")
	}

	return nil
}

// Relational tells whether a column is relational.
func (c *Column) Relational() bool {
	if _, yes := c.Plugin().(ReferenceValidator); yes {
		return yes
	}

	return false
}

// Indexable tells whether index can be created on this column.
func (c *Column) Indexable() bool {
	if _, yes := c.Plugin().(Indexable); yes {
		return yes
	}

	return false
}

// Indexer returns concrete Index type of the column if present.
// Returns nil if the column is not indexed.
func (c *Column) Indexer() Index {
	if !c.Indexed {
		return nil
	}

	idx := c.Plugin().(Indexable).IndexKind()
	// prepare the index
	idx.setUnique(c.Unique)
	idx.setModelID(c.ModelID)
	return idx
}

func (c *Column) createIndexIfNec(val string) (idx Index, err error) {
	if !c.Indexed {
		return
	}

	if p, yes := c.Plugin().(Indexable); yes {
		if idx, err = p.CreateIndex(val); err == nil && idx != nil {
			idx.setUnique(c.Unique)
			idx.setModelID(c.ModelID)
		}

		return
	}

	return
}

// ClobWise tells whether the column cell will be stored in clob table.
func (c *Column) ClobWise() bool {
	_, yes := clobColumns[c.Kind]
	return yes
}

// clobSlot returns slot in Clob for this column.
// e,g. Column.Slot be 601, clobSlot() will return 1.
func (c *Column) clobSlot() int16 {
	if _, yes := clobColumns[c.Kind]; yes {
		return c.Slot - clobSlotBarrier
	}

	return 0
}

// introspect validates the column itself and fix internal state before add/update the column itself.
func (c *Column) introspect() error {
	if c.Deprecated {
		// should never happen: if it ever happens, programmer's bug
		return fmt.Errorf("using a deprecated column:%s not allowed", c.Name)
	}

	if columnNameReserved(c.Name) {
		return fmt.Errorf("column:%s reserved", c.Name)
	}

	// maintain internal state consistency

	if c.Sortable && !c.Indexed {
		c.Indexed = true
	}

	if c.Unique {
		c.Required = true
		c.Indexed = true // will be corrected very soon if it is not indexable
	}

	if c.Relational() || c.Virtual() {
		c.ReadOnly = true
	}

	if _, yes := alwaysIndexColumns[c.Kind]; yes {
		c.Indexed = true
	} else if !c.Indexable() {
		// one-vote veto
		if c.Unique {
			return fmt.Errorf("column:%s wants unique, but its not indexable", c.Name)
		}
		c.Indexed = false
		c.Sortable = false
	}

	// let plugin run if nec
	if p, yes := c.Plugin().(Introspector); yes {
		if err := p.Introspect(); err != nil {
			return err
		}
	}

	return nil
}

// Plugin returns the plugin instance of the column.
func (c *Column) Plugin() ColumnPlugin {
	if c.pluginCache == nil {
		c.pluginCache = columnPlugins[c.Kind](c, c.PluginContext)
	}

	return c.pluginCache
}
