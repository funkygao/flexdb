package localfile

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/agile-app/flexdb/internal/telemetry"
	"github.com/funkygao/go-metrics"
)

type localFile struct {
	reg      metrics.Registry
	interval time.Duration

	quiting, quit chan struct{}
}

// Offer provides the local file telemetry reporter.
func Offer(reg metrics.Registry, interval time.Duration) telemetry.Reporter {
	return &localFile{
		reg:      reg,
		interval: interval,
		quiting:  make(chan struct{}),
		quit:     make(chan struct{}),
	}
}

func (f *localFile) Name() string {
	return "localFile"
}

func (f *localFile) Stop() {
	close(f.quiting)
	<-f.quit
}

func (f *localFile) Start() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	ticker := time.Tick(f.interval)
	for {
		select {
		case <-f.quiting:
			// drain
			close(f.quit)

		case <-ticker:
			f.dump()
		}
	}
}
