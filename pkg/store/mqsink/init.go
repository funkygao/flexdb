package mqsink

import (
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
)

func init() {
	store.RegisterStoreEngine(entity.StoreEngineMQ, func(model *entity.Model) store.DataStore {
		return &mqSink{topic: model.Name}
	})
}
