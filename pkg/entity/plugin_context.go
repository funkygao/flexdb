package entity

import (
	"github.com/agile-app/flexdb/pkg/dto"
)

// PluginContext is provided by runtime engine for column plugins to do more.
type PluginContext interface {

	// PIN returns current login user: Deprecated
	PIN() string

	ModelAccessorOf(modelID int64, preloads ...string) (PluginModelAccessor, error)
}

// PluginModelAccessor is the interface column plugin requires to access model.
type PluginModelAccessor interface {

	// EntityModel return current model.
	EntityModel() *Model

	RetrieveRow(rowID uint64, withChildren bool) (*dto.MasterDetail, error)

	SlotPickList(slot int16) (picklist []*PickItem, err error)
}
