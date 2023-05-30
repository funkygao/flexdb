package entity

import (
	"sort"
	"time"

	"github.com/agile-app/flexdb/pkg/dto"
)

// Model is user defined business model.
// Model is the runtime interpreter sitting between user and the underlying store(PDM).
// Model is a LDM(logical data model) with which users needn't care about PDM.
type Model struct {
	// AppID specifies which app this model belongs to.
	AppID int64 `gorm:"column:app_id" json:"app_id,omitempty"`

	ID     int64  `gorm:"column:id; AUTO_INCREMENT" json:"id,omitempty"`
	Name   string `gorm:"column:name" json:"name,omitempty"`
	Remark string `gorm:"column:remark" json:"remark,omitempty"`

	// IdentName is the primary human recognizable identification column name.
	// Once set, cannot change.
	IdentName string `gorm:"column:ident_name" json:"ident_name,omitempty"`

	Kind        ModelKind `gorm:"column:kind" json:"-"`
	StoreEngine string    `gorm:"column:store" json:"-"`

	// Feature stores features of this model: a bitmap like storage.
	Feature ModelFeature `gorm:"column:feature" json:"-"`

	CTime time.Time `gorm:"column:ctime; type:timestamp; default: NOW();" json:"ctime,omitempty"`
	MTime time.Time `gorm:"column:mtime; type:timestamp; default: NOW();" json:"mtime,omitempty"`
	CUser string    `gorm:"column:cuser" json:"cuser,omitempty"`
	MUser string    `gorm:"column:muser" json:"muser,omitempty"`

	Deleted bool `gorm:"column:deleted" json:"-"`

	// Ver stores version of the model. Each update will increment this value.
	Ver int `gorm:"column:ver" json:"ver"`

	//=============
	// associations
	//=============

	// Slots holds user defined columns of the model, builtin columns exclusive.
	Slots []*Column `gorm:"foreignKey:ModelID" json:"slots,omitempty"`

	// Triggers holds triggers of the model.
	Triggers []ModelTrigger `gorm:"foreignKey:ModelID" json:"-"`

	// Children holds children models of the model.
	Children []Relation `gorm:"foreignKey:ToModelID" json:"-"`

	//==========
	// transient
	//==========

	// H is DTO storage for UI form that transports extra arbitrary data.
	H dto.H `gorm:"-" json:"h,omitempty"`

	OrgID int64 `gorm:"-"`

	slotsCache  map[string]*Column // key is column.name
	sortedSlots Columns            // cache
}

// KindLabel returns model kind in human readable text.
func (m *Model) KindLabel() string {
	return m.Kind.Label()
}

// StoreEngineLabel renders store engine in human readable content.
func (m *Model) StoreEngineLabel() string {
	if m.StoreEngine == "" {
		return StoreEngineFlexDB
	}

	return m.StoreEngine
}

// Virtual tells whether data of the model reside on FlexDB.
func (m *Model) Virtual() bool {
	return m.StoreEngineLabel() != StoreEngineFlexDB
}

// SortedSlots returns sorted slots by ordinal.
func (m *Model) SortedSlots() Columns {
	if m.sortedSlots == nil {
		m.sortedSlots = make(Columns, len(m.Slots))
		for i, c := range m.Slots {
			m.sortedSlots[i] = c
		}

		// sort inside stateless app server instead of RDBMS.
		// select * from field where model_id in (?,?,?) order by ordinal;
		// this SQL will use filesort instead of index sort because IN is range operation.
		sort.Sort(m.sortedSlots)
	}

	return m.sortedSlots
}

// SlotByName returns a slot column by column name.
// Will return nil if not found.
func (m *Model) SlotByName(name string) *Column {
	if s, present := m.slotsCache[name]; present {
		return s
	}

	// loop all slots and cache once for all
	m.slotsCache = make(map[string]*Column, len(m.Slots))
	for _, c := range m.Slots {
		m.slotsCache[c.Name] = c
	}

	return m.slotsCache[name]
}

// SlotBySlotID gets a column by slot id.
func (m *Model) SlotBySlotID(slotID int16) *Column {
	for _, c := range m.Slots {
		if c.Slot == slotID {
			return c
		}
	}

	return nil
}

// SlotByColumnID gets a column by column id.
func (m *Model) SlotByColumnID(columnID int64) *Column {
	for _, c := range m.Slots {
		if c.ID == columnID {
			return c
		}
	}

	return nil
}

// BuiltinColumns return builtin columns of the model.
func (m *Model) BuiltinColumns() []*Column {
	if m.Virtual() {
		return nil
	}

	return builtinColumns
}

// TotalColumnsN returns total columns count of the model: builtin columns + slots
func (m *Model) TotalColumnsN() int {
	return len(m.Slots) + len(builtinColumns)
}

// totalClobSlotsN returns count of slot that will be persisted in RowClob.
func (m *Model) totalClobSlotsN() (n int) {
	for _, c := range m.Slots {
		if c.ClobWise() {
			n++
		}
	}

	return
}

// totalIndexesN returns total count of indexed slots.
func (m *Model) totalIndexesN() (n int) {
	for _, c := range m.Slots {
		if c.Indexed {
			n++
		}
	}

	return
}

// InvokeTriggers invoke all triggers of the model.
func (m *Model) InvokeTriggers(ctx TriggerContext, row *Row, event triggerEvent) error {
	for _, t := range m.Triggers {
		var mappings = map[triggerEvent]func(ctx TriggerContext, m *Model, r *Row) error{
			TriggerBeforeCreate: t.InvokeBeforeInsert,
			TriggerAfterCreate:  t.InvokeAfterInsert,
			TriggerBeforeUpdate: t.InvokeBeforeUpdate,
			TriggerAfterUpdate:  t.InvokeAfterUpdate,
			TriggerBeforeDelete: t.InvokeBeforeDelete,
			TriggerAfterDelete:  t.InvokeAfterDelete,
		}

		if err := mappings[event](ctx, m, row); err != nil {
			return err
		}
	}

	return nil
}
