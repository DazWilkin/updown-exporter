# Prometheus Exporter for [updown.io](https://updown.io)

[![build-containers](https://github.com/DazWilkin/updown-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/updown-exporter/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazWilkin/updown-exporter.svg)](https://pkg.go.dev/github.com/DazWilkin/updown-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/dazwilkin/updown-exporter)](https://goreportcard.com/report/github.com/dazwilkin/updown-exporter)

## Metrics

|Name|Type|Description|
|----|----|-----------|
|`Up`|Counter|Check metrics|

## Image

`ghcr.io/dazwilkin/updown-exporter:ce7fc81da54e017cd62ebfc433fa901b43e60eb0`

## API Key

The Exporter needs access to an updown API Key

```bash
export API_KEY="[YOUR-API-KEY]"
```

## Go

```bash
export API_KEY="[YOUR-API-KEY]"

go run . \
--endpoint=0.0.0.0:8080 \
--path=/metrics
```

## Container

```bash
API_KEY="[YOUR-API-KEY]"

IMAGE="ghcr.io/dazwilkin/updown-exporter:ce7fc81da54e017cd62ebfc433fa901b43e60eb0"

podman run \
--interactive --tty --rm \
--publish=8080:8080 \
--env=API_KEY=${API_KEY} \
${IMAGE} \
  --endpoint=0.0.0.0:8080 \
  --path=/metrics
```

## Prometheus

```YAML
global:
  scrape_interval: 1m
  evaluation_interval: 1m


  # updown Exporter
- job_name: "updown-exporter"
  static_configs:
  - targets:
    - "localhost:8080"
```

<hr/>
<br/>
<a href="https://www.buymeacoffee.com/dazwilkin" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>