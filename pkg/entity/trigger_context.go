package entity

import "github.com/agile-app/flexdb/pkg/dto"

// TriggerContext is a context holder that will make trigger able to do more!
// With TriggerContext design, triggers can be decoupled from store.
type TriggerContext interface {
	LoadModel(modelID int64) (TriggerModelAccessor, error)
}

// TriggerModelAccessor is the interface trigger requires to access model.
type TriggerModelAccessor interface {

	// EntityModel return current model.
	EntityModel() *Model

	// RetrieveRow retrieves a row from the curent model by row id.
	RetrieveRow(rowID uint64, withChildren bool) (*dto.MasterDetail, error)
}
