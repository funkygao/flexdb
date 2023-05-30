package upstream

import (
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
)

func init() {
	store.RegisterStoreEngine(entity.StoreEngineUpstream, func(model *entity.Model) store.DataStore {
		return &upstream{Model: model}
	})
}
