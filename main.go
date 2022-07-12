package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	stdlog "log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/DazWilkin/updown-exporter/collector"
	"github.com/DazWilkin/updown-exporter/updown"

	"github.com/go-logr/stdr"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace string = "updown"
	subsystem string = "exporter"
	version   string = "v0.0.1"
)
const (
	rootTemplate string = `
{{- define "content" }}
<!DOCTYPE html>
<html lang="en-US">
<head>
<title>Prometheus Exporter for updown</title>
<style>
body {
  font-family: Verdana;
}
</style>
</head>
<body>
	<h2>Prometheus Exporter for updown</h2>
	<hr/>
	<ul>
	<li><a href="{{ .MetricsPath }}">metrics</a></li>
	<li><a href="/healthz">healthz</a></li>
	</ul>
</body>
</html>
{{- end}}
`
)

var (
	// GitCommit is the git commit value and is expected to be set during build
	GitCommit string
	// GoVersion is the Golang runtime version
	GoVersion = runtime.Version()
	// OSVersion is the OS version (uname --kernel-release) and is expected to be set during build
	OSVersion string
	// StartTime is the start time of the exporter represented as a UNIX epoch
	StartTime = time.Now().Unix()
)
var (
	endpoint    = flag.String("endpoint", "0.0.0.0:8080", "The endpoint of the HTTP server")
	metricsPath = flag.String("path", "/metrics", "The path on which Prometheus metrics will be served")
)
var (
	name string = fmt.Sprintf("%s_%s", namespace, subsystem)
)

type Content struct {
	MetricsPath string
}

func handleHealthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func handleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t := template.Must(template.New("content").Parse(rootTemplate))
	if err := t.ExecuteTemplate(w, "content", Content{MetricsPath: *metricsPath}); err != nil {
		log.Fatal("unable to execute template")
	}
}

func main() {
	log := stdr.NewWithOptions(stdlog.New(os.Stderr, "", stdlog.LstdFlags), stdr.Options{LogCaller: stdr.All})
	log = log.WithName("main")

	flag.Parse()
	if *endpoint == "" {
		log.Info("Expected flag `--endpoint`")
		os.Exit(1)
	}

	var key string
	if key = os.Getenv("API_KEY"); key == "" {
		log.Info("Expected `API_KEY` in the environment")
		os.Exit(1)
	}

	if GitCommit == "" {
		log.Info("GitCommit value unchanged: expected to be set during build")
	}
	if OSVersion == "" {
		log.Info("OSVersion value unchanged: expected to be set during build")
	}

	// Objects that holds GCP-specific resources (e.g. projects)
	client := updown.NewClient(key, log)

	registry := prometheus.NewRegistry()

	s := collector.System{
		Namespace: namespace,
		Subsystem: subsystem,
		Version:   version,
	}

	b := collector.Build{
		OsVersion: OSVersion,
		GoVersion: GoVersion,
		GitCommit: GitCommit,
		StartTime: StartTime,
	}
	registry.MustRegister(collector.NewExporterCollector(s, b, log))
	registry.MustRegister(collector.NewChecksCollector(s, client, log))
	registry.MustRegister(collector.NewMetricsCollector(s, client, log))

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleRoot))
	mux.Handle("/healthz", http.HandlerFunc(handleHealthz))
	mux.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	log.Info("Starting",
		"endpoint", *endpoint,
		"metrics", *metricsPath,
	)
	log.Error(http.ListenAndServe(*endpoint, mux), "unable to start server")
}
