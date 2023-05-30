package store

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"gorm.io/gorm"
)

// store is not goroutine safe: its singleton.
type store struct {
	metaSource  *gorm.DB
	dataSources []*gorm.DB
}

// Offer creates Store with MySQL as underlying storage engine.
func Offer(metaSource *gorm.DB, dataSources []*gorm.DB) Store {
	return &store{metaSource: metaSource, dataSources: dataSources}
}

func (s *store) InferOrgStore(c context.RESTContext) OrgStore {
	return newOrgStore(s, c)
}

func (s store) InferAppStore(c context.RESTContext) AppStore {
	return s.InferOrgStore(c).AppStoreOf(c.AppID())
}

// mdb returns the singleton meta DB.
func (s *store) mdb() *gorm.DB {
	return s.metaSource
}

// ddb returns a data DB.
func (s *store) ddb(appID int64) *gorm.DB {
	return s.dataSources[0] // TODO sharding policy
}
