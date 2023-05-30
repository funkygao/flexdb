// Package telemetry persists github.com/funkygao/go-metrics
// metrics.Registry to durable storage.
package telemetry

// A Reporter continuously scans metrics.Registry and
// persists all metrics to durable storage.
type Reporter interface {
	Name() string

	Start()
	Stop()
}

// Provider is the default Reporter implementation.
var Provider Reporter
