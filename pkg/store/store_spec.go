package store

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
)

// Store manages underlying storage engine for FlexDB.
type Store interface {

	// InferOrgStore infers org store from RESTContext.
	InferOrgStore(context.RESTContext) OrgStore

	// InferAppStore infers app store from RESTContext.
	InferAppStore(context.RESTContext) AppStore
}

// OrgStore is store for the org.
type OrgStore interface {

	// AppStoreOf creates a AppStore from an app id.
	AppStoreOf(appID int64) AppStore

	// LoadApp loads an app.
	// If preloads not empty, will eager load correstponding entities.
	LoadApp(appID int64, preloads ...string) (app *entity.App, err error)

	// CreateApp persists the app for the current org.
	CreateApp(app *entity.App) (err error)

	// UpdateApp updates an existing app.
	UpdateApp(app *entity.App) (err error)

	// FindApps find apps of the current org.
	FindApps(criteria dto.Criteria, page, pageSize int) (apps []*entity.App, err error)

	ShareApp(app *entity.App, shareTo *entity.Share) (err error)
}

// AppStore is store for App.
type AppStore interface {

	// ModelStoreOf create a ModelStoreOf based on the loaded model.
	ModelStoreOf(model *entity.Model) ModelStore

	// CreateModel persists the model itself for the current app.
	CreateModel(model *entity.Model) error

	UpdateModel(model *entity.Model) error

	// Models returns all models of the current app.
	// If preloads not empty, will eager load correstponding entities.
	LoadModels(preloads ...string) (models []entity.Model, err error)

	// LoadModel loads specified model together with its columns and return the store repo.
	// If preloads not empty, will eager load correstponding entities.
	LoadModel(modelID int64, preloads ...string) (ModelStore, error)
}

// ModelStore is store for Model.
// ModelStore manages all metadata and treats everything as Table: like facebook presto.
type ModelStore interface {
	entity.PluginModelAccessor

	// DS returns the actual data source of the model.
	DS() DataStore

	AddColumns(cols []*entity.Column) error
	DeprecateColumn(*entity.Column) error
	UpdateColumn(column *entity.Column) error
	ReorderColumns(cols []entity.Column) error

	CreateSlotPickItem(slot int16, val string) (*entity.PickItem, error)
	UpdateSlotPickItem(id int64, val string) error
	SlotPickList(slot int16) (picklist []*entity.PickItem, err error)
}

// DataStore is store for data.
// DataStore is pluggable, like MySQL storage engine design.
type DataStore interface {
	CreateRow(rd dto.RowData) (rowID uint64, err error)
	RetrieveRow(rowID uint64, withChildren bool) (*dto.MasterDetail, error)
	FindRows(selectFields []string, criteria dto.Criteria, page, pageSize int) ([]dto.RowData, error)
	UpdateRow(rd dto.RowData) error
	QuickSave(qs dto.QuickSave) error
	DeleteRow(id uint64) error
}

// Provider provides the default Store instance.
var Provider Store
