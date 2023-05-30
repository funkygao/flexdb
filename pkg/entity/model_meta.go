package entity

import (
	"fmt"
)

// AddSlots refines columns before adding slots to the model.
// AddSlots will not change state of model but toAddSlots items will be updated.
func (m *Model) AddSlots(toAddSlots []*Column, pc PluginContext) (err error) {
	// constraint: max slot
	if len(m.Slots)+len(toAddSlots) > maxSlots {
		return fmt.Errorf("one model can have at most %d slots", maxSlots)
	}

	// in physical column table, (model_id, column_name) is unique index
	// that can completely avoid dup column case

	currentSlot := int16(len(m.Slots)) + 1
	currentClobSlot := int16(m.totalClobSlotsN()) + clobSlotBarrier + 1
	toAddIndexN, toAddClobSlotN := 0, 0
	for _, c := range toAddSlots {
		c.ModelID = m.ID
		c.PluginContext = pc

		if err = c.introspect(); err != nil {
			return err
		}

		if c.Indexed {
			toAddIndexN++
		}

		if c.Slot == 0 {
			// slot not specified: auto assign slot
			if c.ClobWise() {
				c.Slot = currentClobSlot
				currentClobSlot++
				toAddClobSlotN++
			} else {
				c.Slot = currentSlot
				currentSlot++
			}
		}

		if p, yes := c.Plugin().(ReferenceValidator); yes {
			if err = p.ValidateReference(m, toAddSlots); err != nil {
				return err
			}
		}
	}

	// constraint: max indexes per model
	if toAddIndexN+m.totalIndexesN() > maxIndexesPerModel {
		return fmt.Errorf("one model can have at most %d indexes", maxIndexesPerModel)
	}

	// constraint: max clob columns
	if toAddClobSlotN+m.totalClobSlotsN() > maxClobSlots {
		return fmt.Errorf("one model can have at most %d clob column", maxClobSlots)
	}

	return
}

// UpdateSlot updates an existing column of the model.
func (m *Model) UpdateSlot(slot *Column, pc PluginContext) (err error) {
	if err = slot.introspect(); err != nil {
		return err
	}

	// TODO more constraints

	return
}
