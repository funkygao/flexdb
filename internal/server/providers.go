package server

import (
	"time"

	"github.com/agile-app/flexdb/internal/cache"
	"github.com/agile-app/flexdb/internal/cache/blackhole"
	"github.com/agile-app/flexdb/internal/cache/gcache"
	"github.com/agile-app/flexdb/internal/sla"
	"github.com/agile-app/flexdb/internal/sla/dummy"
	"github.com/agile-app/flexdb/internal/telemetry"
	"github.com/agile-app/flexdb/internal/telemetry/localfile"
	"github.com/agile-app/flexdb/pkg/driver/mysql"
	"github.com/agile-app/flexdb/pkg/driver/postgres"
	"github.com/funkygao/go-metrics"
	"gorm.io/gorm"
)

var (
	dataSourceDrivers = map[string]func() (*gorm.DB, []*gorm.DB, error){
		"mysql":    mysql.Offer,
		"postgres": postgres.Offer,
	}

	slaProviders = map[string]func() sla.SLA{
		"dummy": dummy.Offer,
	}

	cacheProviders = map[string]func() cache.Cache{
		"blackhole": blackhole.Offer,
		"gcache":    gcache.Offer,
	}

	telemetryProviders = map[string]func(metrics.Registry, time.Duration) telemetry.Reporter{
		"localFile": localfile.Offer,
	}
)
