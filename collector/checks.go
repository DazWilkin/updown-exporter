package collector

import (
	"strconv"
	"sync"

	"github.com/DazWilkin/updown-exporter/updown"
	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
)

// ChecksCollector is a type that represents updown Checks
type ChecksCollector struct {
	System  System
	Client  *updown.Client
	Log     logr.Logger
	Enabled *prometheus.Desc
}

// NewChecksCollector is a function that returns a new ChecksCollector
func NewChecksCollector(s System, client *updown.Client, log logr.Logger) *ChecksCollector {
	subsystem := "checks"
	return &ChecksCollector{
		System: s,
		Client: client,
		Log:    log,
		Enabled: prometheus.NewDesc(
			prometheus.BuildFQName(s.Namespace, subsystem, "enabled"),
			"status of check (enabled=1)",
			[]string{
				"token",
				"url",
				"status",
				"ssl_valid",
			},
			nil,
		),
	}
}

// Collect implements Prometheus' Collector interface and is used to collect metrics
func (c *ChecksCollector) Collect(ch chan<- prometheus.Metric) {
	log := c.Log.WithName("Collect")

	checks, err := c.Client.GetChecks()
	if err != nil {
		log.Info("Unable to get Checks")
		return
	}

	var wg sync.WaitGroup
	for _, check := range checks {
		wg.Add(1)
		go func(check updown.Check) {
			defer wg.Done()
			ch <- prometheus.MustNewConstMetric(
				c.Enabled,
				prometheus.CounterValue,
				func(enabled bool) (result float64) {
					if enabled {
						result = 1.0
					}
					return result
				}(check.Enabled),
				[]string{
					check.Token,
					check.URL,
					strconv.FormatUint(uint64(check.LastStatus), 10),
					strconv.FormatBool(check.SSL.Valid),
				}...,
			)
		}(check)
	}
	wg.Wait()
}

// Describe implements Prometheus' Collector interface is used to describe metrics
func (c *ChecksCollector) Describe(ch chan<- *prometheus.Desc) {
	// log := c.Log.WithName("Describe")
	ch <- c.Enabled
}
