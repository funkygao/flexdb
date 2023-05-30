package store

import (
	"github.com/agile-app/flexdb/internal/cache"
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// appStore is goroutine safe.
type appStore struct {
	store *store

	c context.RESTContext

	orgID int64
	appID int64

	metaDB *gorm.DB

	modelCache map[int64]ModelStore
}

var (
	_ entity.PluginContext = (*appStore)(nil) // appStore is a PluginContext
)

func newAppStore(store *store, orgID, appID int64, c context.RESTContext) AppStore {
	return &appStore{
		store:      store,
		metaDB:     store.metaSource,
		orgID:      orgID,
		appID:      appID,
		c:          c,
		modelCache: make(map[int64]ModelStore),
	}
}

func (a *appStore) ModelStoreOf(m *entity.Model) ModelStore {
	return newModelStore(a, m)
}

func (a *appStore) PIN() string {
	return a.c.PIN()
}

// CreateModel persists the model. ID will be set for the model.
func (a *appStore) CreateModel(model *entity.Model) error {
	model.AppID = a.appID
	err := a.metaDB.Omit(clause.Associations).Create(model).Error
	if err == nil {
		cache.Provider.Evict(appModelsHintID, a.appID)
	}
	return err
}

func (a *appStore) UpdateModel(model *entity.Model) error {
	updates := map[string]interface{}{
		"muser":   model.MUser,
		"mtime":   model.MTime,
		"remark":  model.Remark,
		"feature": model.Feature,
	}
	if err := a.metaDB.Omit(clause.Associations).Model(model).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (a *appStore) LoadModels(preloads ...string) (models []entity.Model, err error) {
	if v, err := cache.Provider.Get(appModelsHintID, a.appID); err == nil {
		return v.([]entity.Model), nil
	}

	db := a.metaDB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	err = db.Order("id asc").Where("app_id=?", a.appID).Find(&models).Error
	if err == nil {
		cache.Provider.SetWithExpire(appModelsHintID, a.appID, models, appModelsTTL)
	}
	return
}

func (a *appStore) ModelAccessorOf(modelID int64, preloads ...string) (entity.PluginModelAccessor, error) {
	return a.LoadModel(modelID, preloads...)
}

func (a *appStore) LoadModel(modelID int64, preloads ...string) (ModelStore, error) {
	// session based cache
	if m, present := a.modelCache[modelID]; present {
		// bingo!
		return m, nil
	}

	// instance based cache
	if ms, err := cache.Provider.Get(modelStoreCacheHintID, modelID); err == nil {
		// bingo!
		return ms.(ModelStore), nil
	}

	var model entity.Model
	model.AppID = a.appID
	db := a.metaDB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	if err := db.Take(&model, modelID).Error; err != nil {
		return nil, err
	}

	model.OrgID = a.orgID
	m := newModelStore(a, &model)
	a.modelCache[modelID] = m

	for _, slot := range model.Slots {
		slot.PluginContext = a
	}

	cache.Provider.SetWithExpire(modelStoreCacheHintID, modelID, m, modelStoreCacheTTL)
	return m, nil
}
