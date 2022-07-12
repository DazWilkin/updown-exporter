package collector

import (
	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
)

// ExporterCollector collects metrics, mostly runtime, about this exporter in general.
type ExporterCollector struct {
	System System
	Build  Build
	Log    logr.Logger

	StartTime *prometheus.Desc
	BuildInfo *prometheus.Desc
}

// Build is a type that represents the build info
type Build struct {
	GitCommit string
	GoVersion string
	OsVersion string
	StartTime int64
}

// NewExporterCollector returns a new ExporterCollector.
func NewExporterCollector(s System, b Build, log logr.Logger) *ExporterCollector {
	return &ExporterCollector{
		System: s,
		Build:  b,
		Log:    log,
		StartTime: prometheus.NewDesc(
			prometheus.BuildFQName(s.Namespace, s.Subsystem, "start_time"),
			"Exporter start time in Unix epoch seconds",
			nil,
			nil,
		),
		BuildInfo: prometheus.NewDesc(
			prometheus.BuildFQName(s.Namespace, s.Subsystem, "build_info"),
			"A metric with a constant '1' value labeled by OS version, Go version, and the Git commit of the exporter",
			[]string{"os_version", "go_version", "git_commit"},
			nil,
		),
	}
}

// Collect implements Prometheus' Collector interface and is used to collect metrics
func (c *ExporterCollector) Collect(ch chan<- prometheus.Metric) {
	log := c.Log.WithName("Collect")
	log.Info("Metrics",
		"start_time", c.Build.StartTime,
	)
	ch <- prometheus.MustNewConstMetric(
		c.StartTime,
		prometheus.GaugeValue,
		float64(c.Build.StartTime),
	)
	ch <- prometheus.MustNewConstMetric(
		c.BuildInfo,
		prometheus.CounterValue,
		1.0,
		c.Build.OsVersion, c.Build.GoVersion, c.Build.GitCommit,
	)
}

// Describe implements Prometheus' Collector interface and is used to describe metrics
func (c *ExporterCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.StartTime
}
