package perm

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func Offer(db *gorm.DB) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}

	if adapter == nil {

	}
}
