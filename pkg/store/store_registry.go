package store

import (
	"fmt"

	"github.com/agile-app/flexdb/pkg/entity"
)

type dataStoreFactory func(model *entity.Model) DataStore

var (
	storeEngines = make(map[string]dataStoreFactory, 2)
)

// RegisterStoreEngine registers a data store engine factory.
func RegisterStoreEngine(name string, factory dataStoreFactory) {
	if _, present := storeEngines[name]; present {
		panic(fmt.Errorf("store engine:%s already registered", name))
	}

	storeEngines[name] = factory
}
