package entity

import (
	"fmt"
)

var (
	// kind:factory
	columnPlugins = make(map[ColumnKind]func(*Column, PluginContext) ColumnPlugin, 20)

	reservedColumnMap = make(map[int]struct{})

	// trigger_id:spec
	triggers = make(map[int]TriggerSpec)
)

// RegisterTrigger registers a trigger by id.
func RegisterTrigger(t TriggerSpec) {
	if _, present := triggers[t.ID()]; present {
		panic(fmt.Sprintf("trigger[%v] cannot register twice", t.ID()))
	}

	triggers[t.ID()] = t
}

// LocateTrigger locates a trigger by id and return its presence.
func LocateTrigger(triggerID int) (TriggerSpec, bool) {
	t, present := triggers[triggerID]
	return t, present
}

// RegisterColumnPlugin registers a column plugin.
func RegisterColumnPlugin(kind ColumnKind, factory func(*Column, PluginContext) ColumnPlugin, reservedColumnSlots ...int) {
	if _, present := columnPlugins[kind]; present {
		panic(fmt.Sprintf("ColumnPlugin[%v] cannot register twice", kind))
	}

	if len(reservedColumnSlots) > 0 {
		for _, n := range reservedColumnSlots {
			if _, present := reservedColumnMap[n]; present {
				panic(fmt.Errorf("reserved column:%d cannot reserve twice", n))
			}

			reservedColumnMap[n] = struct{}{}
		}
	}

	columnPlugins[kind] = factory
}
