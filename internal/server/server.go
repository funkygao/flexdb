package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/agile-app/flexdb/internal/cache"
	"github.com/agile-app/flexdb/internal/controller"
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/internal/controller/context/amis"
	"github.com/agile-app/flexdb/internal/controller/data"
	"github.com/agile-app/flexdb/internal/controller/meta"
	"github.com/agile-app/flexdb/internal/controller/schema"
	"github.com/agile-app/flexdb/internal/controller/user"
	"github.com/agile-app/flexdb/internal/i18n"
	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/internal/router"
	"github.com/agile-app/flexdb/internal/sla"
	"github.com/agile-app/flexdb/internal/telemetry"
	"github.com/agile-app/flexdb/internal/worker"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/funkygao/go-metrics"
	"github.com/funkygao/log4go"
	"github.com/gin-gonic/gin"

	// data sources
	_ "github.com/agile-app/flexdb/pkg/store/mock"
	_ "github.com/agile-app/flexdb/pkg/store/mqsink"
	_ "github.com/agile-app/flexdb/pkg/store/upstream"
	_ "github.com/agile-app/flexdb/pkg/store/usf"

	// plugins
	_ "github.com/agile-app/flexdb/plugins"
)

// Server is a server instance that runs on a Linux box.
// Server orchestrates components of FlexDB and serves the world.
type Server struct {
	router *gin.Engine

	wg         sync.WaitGroup
	shutdownCh <-chan struct{}

	mc controller.MetaHandler
	dc controller.DataHandler
	sc controller.SchemaHandler
	uc controller.UserHandler
}

// New creates a Server.
func New(shutdownCh <-chan struct{}) *Server {
	return &Server{
		router:     router.Offer(),
		mc:         meta.Offer(),
		dc:         data.Offer(),
		sc:         schema.Offer(),
		uc:         user.Offer(),
		shutdownCh: shutdownCh,
	}
}

// Prepare prepares all works before start to serve the world.
func (Server) Prepare() {
	i18n.B = i18n.Offer(profile.P.DefaultLocale)

	// REST context
	context.Offer = func(c *gin.Context) context.RESTContext {
		return amis.Of(c)
	}

	cf := profile.P
	metaSource, dataSources, err := dataSourceDrivers[cf.Driver]()
	if err != nil {
		panic(err)
	}

	store.Provider = store.Offer(metaSource, dataSources)

	// SLA provider
	sla.Provider = slaProviders[cf.SLA]()

	// cache provider
	cache.Provider = cacheProviders[cf.Cache]()

	// telemetry
	telemetry.Provider = telemetryProviders[cf.Telemetry](metrics.DefaultRegistry, 10*time.Minute)
	go func() {
		log4go.Info("starting telemetry:%s", telemetry.Provider.Name())
		telemetry.Provider.Start()
	}()

	entity.Prepare()

	// workers
	worker.StartAll(nil)

	// connectors init
}

// ServeForever starts the backend server and serves request forever.
func (s *Server) ServeForever() (err error) {
	httpServer := &http.Server{
		Addr:           profile.P.BindAddr,
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		IdleTimeout:    5 * time.Minute,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 16 << 10,
	}
	return httpServer.ListenAndServe()
}
