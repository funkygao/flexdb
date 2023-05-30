package localfile

import (
	"github.com/funkygao/go-metrics"
	"github.com/funkygao/log4go"
)

func (f *localFile) dump() {
	f.reg.Each(func(name string, i interface{}) {
		switch m := i.(type) {
		case metrics.Counter:
			log4go.Info("%s %d", name, m.Count())

		case metrics.Gauge:
			log4go.Info("%s %d", name, m.Value())

		case metrics.GaugeFloat64:
			log4go.Info("%s %d", name, m.Value())

		case metrics.Histogram:
			ps := m.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			item := map[string]interface{}{
				"N":    m.Count(),
				"max":  m.Max(),
				"mean": m.Mean(),
				"min":  m.Min(),
				"p50":  ps[0],
				"p75":  ps[1],
				"p95":  ps[2],
				"p99":  ps[3],
				//"p999": ps[4],
				//"stddev": m.StdDev(),
				//"variance": m.Variance(),
			}
			log4go.Info("%s %v", name, item)

		case metrics.Meter:
			item := map[string]interface{}{
				"N":    m.Count(),
				"m1":   m.Rate1(),
				"m5":   m.Rate5(),
				"m15":  m.Rate15(),
				"mean": m.RateMean(),
			}
			log4go.Info("%s %v", name, item)

		}
	})
}
