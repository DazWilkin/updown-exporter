# Prometheus Exporter for [updown.io](https://updown.io)

[![build-containers](https://github.com/DazWilkin/updown-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/updown-exporter/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazWilkin/updown-exporter.svg)](https://pkg.go.dev/github.com/DazWilkin/updown-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/dazwilkin/updown-exporter)](https://goreportcard.com/report/github.com/dazwilkin/updown-exporter)

## Metrics

Metrics names are prefixed `updown_`.

|Name|Type|Description|
|----|----|-----------|
|`checks_enabled`|Counter|Status of Check (enabled=1)|
|`exporter_build_info`|Counter|Exporter build info|
|`metrics_response_times`|Histogram|Histogram of a Check's response times|

## Image

`ghcr.io/dazwilkin/updown-exporter:ded7fddc410e3c3037fd147609f3b8d5d4aeb4cb`

## [Sigstore](https://sigstore.dev)

`updown-exporter` container images are being signed by Sigstore and may be verified:

```bash
cosign verify \
--key=./cosign.pub \
ghcr.io/dazwilkin/updown-exporter:ded7fddc410e3c3037fd147609f3b8d5d4aeb4cb
```

NOTE `cosign.pub` may be downloaded [here](./cosign.pub)

To install `cosign`, e.g.:

```bash
go install github.com/sigstore/cosign/cmd/cosign@latest
```

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

IMAGE="ghcr.io/dazwilkin/updown-exporter:ded7fddc410e3c3037fd147609f3b8d5d4aeb4cb"

podman run \
--interactive --tty --rm \
--publish=8080:8080 \
--env=API_KEY=${API_KEY} \
${IMAGE} \
  --endpoint=0.0.0.0:8080 \
  --path=/metrics
```

Then browse `http://localhost:8080/metrics` to view the metrics.

## Prometheus

`prometheus.yml`:
```YAML
global:
  scrape_interval: 1m
  evaluation_interval: 1m

scrape_configs:
  # updown Exporter
- job_name: "updown-exporter"
  static_configs:
  - targets:
    - "localhost:8080"
```

## Docker

```bash
API_KEY="[YOUR-API-KEY]"

IMAGE="ghcr.io/dazwilkin/updown-exporter:ded7fddc410e3c3037fd147609f3b8d5d4aeb4cb"

docker run \
--detach --tty --rm \
--name="updown-exporter" \
--env=API_KEY=${API_KEY} \
--publish=8080:8080/tcp \
${IMAGE} \
  --endpoint=0.0.0.0:8080 \
  --path=/metrics

docker run \
--detach --rm --tty \
--name="prometheus" \
--publish=9090:9090/tcp \
--volume=${PWD}/prometheus.yml:/etc/prometheus/prometheus.yml:ro \
docker.io/prom/prometheus:v2.36.2 \
--config.file=/etc/prometheus/prometheus.yml \
--web.enable-lifecycle
```

Then browse:

+ [Exporter](http://localhost:8080/metrics)
+ [Prometheus](http://localhost:9090/targets)

## Podman

```bash
API_KEY="[YOUR-API-KEY]"

IMAGE="ghcr.io/dazwilkin/updown-exporter:ded7fddc410e3c3037fd147609f3b8d5d4aeb4cb"

POD="updown-exporter"

podman pod create \
--name=${POD} \
--publish=5555:8080/tcp \
--publish=9090:9090/tcp

podman run \
--interactive --tty --rm \
--pod=${POD} \
--name="updown-exporter" \
--env=API_KEY=${API_KEY} \
${IMAGE} \
  --endpoint=0.0.0.0:8080 \
  --path=/metrics

podman run \
--detach --rm --tty \
--pod=${POD} \
--name="prometheus" \
--volume=${PWD}/prometheus.yml:/etc/prometheus/prometheus.yml:ro \
docker.io/prom/prometheus:v2.36.2 \
--config.file=/etc/prometheus/prometheus.yml \
--web.enable-lifecycle
```

Then browse:

+ [Exporter](http://localhost:8080/metrics)
+ [Prometheus](http://localhost:9090/targets)

## Raspberry Pi

```bash
if [ "$(getconf LONG_BIT)" -eq 64 ]
then
  # 64-bit Raspian
  ARCH="GOARCH=arm64"
  TAG="arm64"
else
  # 32-bit Raspian
  ARCH="GOARCH=arm GOARM=7"
  TAG="arm32v7"
fi

podman build \
--build-arg=GOLANG_OPTIONS="CGO_ENABLED=0 GOOS=linux ${ARCH}" \
--build-arg=COMMIT=$(git rev-parse HEAD) \
--build-arg=VERSION=$(uname --kernel-release) \
--tag=ghcr.io/dazwilkin/updown-exporter:${TAG} \
--file=./Dockerfile \
.
```

## Similar Exporters

+ [Prometheus Exporter for Azure](https://github.com/DazWilkin/azure-exporter)
+ [Prometheus Exporter for crt.sh](https://github.com/DazWilkin/crtsh-exporter)
+ [Prometheus Exporter for Fly.io](https://github.com/DazWilkin/fly-exporter)
+ [Prometheus Exporter for GoatCounter](https://github.com/DazWilkin/goatcounter-exporter)
+ [Prometheus Exporter for Google Cloud](https://github.com/DazWilkin/gcp-exporter)
+ [Prometheus Exporter for Koyeb](https://github.com/DazWilkin/koyeb-exporter)
+ [Prometheus Exporter for Linode](https://github.com/DazWilkin/linode-exporter)
+ [Prometheus Exporter for PorkBun](https://github.com/DazWilkin/porkbun-exporter)
+ [Prometheus Exporter for updown.io](https://github.com/DazWilkin/updown-exporter)
+ [Prometheus Exporter for Vultr](https://github.com/DazWilkin/vultr-exporter)

<hr/>
<br/>
<a href="https://www.buymeacoffee.com/dazwilkin" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>
