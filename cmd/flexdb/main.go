package main

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/agile-app/flexdb/internal/diagnostics"
	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/internal/server"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/funkygao/golib/daemon"
	"github.com/funkygao/golib/signal"
	"github.com/funkygao/golib/version"
	"github.com/funkygao/log4go"
)

var (
	shutdownOnce sync.Once
	shutdownCh   = make(chan struct{})
)

func main() {
	parseFlags()
	setupLogging()

	if options.printVersion {
		fmt.Printf("FlexDB(%s-%s) built at %s\n", version.Version, version.Revision, version.BuildDate)
		fmt.Printf("built with %s\n", version.GoVersion)
		os.Exit(0)
	}

	if err := profile.LoadFrom(configFn); err != nil {
		panic(err)
	}

	if options.migrateDB {
		db, err := sql.Open("mysql", profile.P.MetaDSN)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		store.MigrateDB(db)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	t0 := time.Now()

	s := server.New(shutdownCh)
	s.PublishAPIs()

	if options.showAPIsAndExit {
		showAPIs()
		os.Exit(0)
	}

	// will run as daemon

	if runtime.GOOS == "linux" && !options.showAPIsAndExit {
		daemon.EnsureServerUlimit()
	}

	fmt.Fprintln(os.Stderr, banner)

	log4go.Info("pprof started at %s", profile.P.PprofEndpoint)
	diagnostics.Start(profile.P.PprofEndpoint)
	go func() {
		log4go.Error(<-diagnostics.Errors)
	}()

	// TODO graceful shutdown
	if false {
		signal.RegisterHandler(func(sig os.Signal) {
			shutdownOnce.Do(func() {
				log4go.Info("FlexDB(%s-%s) received signal: %s, start graceful shutdown...", version.Version, version.Revision,
					strings.ToUpper(sig.String()))

				close(shutdownCh)
			})
		}, /*syscall.SIGINT, */ syscall.SIGTERM) // yes we ignore HUP
	}

	s.Prepare()
	log4go.Trace("%+v", *profile.P)
	log4go.Info("FlexDB(%s-%s) started", version.Version, version.Revision)
	if err := s.ServeForever(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	log4go.Info("%s-%s, %s, bye!", version.Version, version.Revision, time.Since(t0))
	log4go.Close() // flush
}
