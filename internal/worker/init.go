package worker

import (
	"github.com/funkygao/log4go"
	"gorm.io/gorm"
)

var (
	workers = []Worker{}
)

func init() {
	for _, w := range workers {
		w.Init()
	}
}

// StartAll starts all registered workers in async mode.
func StartAll(db *gorm.DB) {
	for _, w := range workers {
		log4go.Info("starting worker:%s", w.Name())
		go w.Run(db)
	}
}
