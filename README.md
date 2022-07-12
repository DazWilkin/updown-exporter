# Prometheus Exporter for [updown.io](https://updown.io)

## Metrics

|Name|Type|Description|
|----|----|-----------|
|`Up`|Counter||

## Image

ghcr.io/dazwilkin/updown-exporter:123456789abcdef0123456789abcdef012345678

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

IMAGE="ghcr.io/dazwilkin/updown-exporter:123456789abcdef0123456789abcdef012345678"

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
