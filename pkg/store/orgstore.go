package store

import (
	"fmt"
	"time"

	"github.com/agile-app/flexdb/internal/cache"
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"gorm.io/gorm/clause"
)

// orgStore is goroutine safe.
type orgStore struct {
	store *store

	c context.RESTContext

	orgID int64 // cache
}

func newOrgStore(store *store, c context.RESTContext) OrgStore {
	return &orgStore{store: store, c: c, orgID: c.OrgID()}
}

func (o *orgStore) AppStoreOf(appID int64) AppStore {
	return newAppStore(o.store, o.orgID, appID, o.c)
}

func (o *orgStore) CreateApp(app *entity.App) (err error) {
	app.OrgID = o.orgID // assign org id
	now := time.Now()
	app.CTime = now
	app.MTime = now
	app.CUser = o.c.PIN()
	app.Status = entity.AppStatusInit
	err = o.store.mdb().Omit(clause.Associations).Create(app).Error
	return
}

func (o *orgStore) UpdateApp(app *entity.App) (err error) {
	updates := map[string]interface{}{
		"muser":       app.MUser, // TODO should be from context?
		"mtime":       app.MTime,
		"description": app.Description,
		"name":        app.Name,
		"visibility":  app.Visibility,
	}
	if app.Logo != "" {
		updates["logo"] = app.Logo
	}
	err = o.store.mdb().Omit(clause.Associations).Model(app).Updates(updates).Error
	return
}

func (o *orgStore) LoadApp(appID int64, preloads ...string) (app *entity.App, err error) {
	if a, err := cache.Provider.Get(appCacheHintID, appID); err == nil {
		return a.(*entity.App), nil
	}

	app = &entity.App{ID: appID}
	db := o.store.mdb()
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	err = db.Take(&app).Error
	if err == nil {
		cache.Provider.SetWithExpire(appCacheHintID, appID, app, appCacheTTL)
	}
	return
}

func (o *orgStore) FindApps(cond dto.Criteria, page, pageSize int) (apps []*entity.App, err error) {
	db := o.store.mdb()
	if profile.Debug() {
		db = db.Debug()
	}
	db = db.Order("org_id,name"). // idx: (org_id, name)
					Where("org_id=?", o.orgID)
	if cond.Size() > 0 {
		for _, c := range cond {
			db = db.Where(fmt.Sprintf("%s %s ?", c.Key, c.Op), c.Val)
		}
	}
	err = db.Find(&apps).Error
	return
}

func (o *orgStore) ShareApp(app *entity.App, shareTo *entity.Share) (err error) {
	err = o.store.mdb().Omit(clause.Associations).Create(shareTo).Error
	return
}
