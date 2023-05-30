package store

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/entity"
)

var _ entity.PluginModelAccessor = (*modelStore)(nil)

// modelStore is goroutine safe.
type modelStore struct {
	*entity.Model

	store *store

	orgID int64

	c  context.RESTContext
	as *appStore
}

func newModelStore(as *appStore, model *entity.Model) *modelStore {
	return &modelStore{
		as:    as,
		c:     as.c,
		store: as.store,
		orgID: as.orgID,
		Model: model,
	}
}

func (m *modelStore) EntityModel() *entity.Model {
	return m.Model
}

func (m *modelStore) DS() DataStore {
	if factory, present := storeEngines[m.StoreEngine]; present {
		return factory(m.Model)
	}

	return m
}
