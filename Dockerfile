ARG GOLANG_VERSION=1.24.3

ARG TARGETOS
ARG TARGETARCH

ARG COMMIT
ARG VERSION

FROM --platform=${TARGETARCH} docker.io/golang:${GOLANG_VERSION} AS build

WORKDIR /updown-exporter

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go .
COPY collector collector
COPY updown updown

ARG TARGETOS
ARG TARGETARCH

ARG COMMIT=""
ARG VERSION=""

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build \
    -ldflags "-X main.OSVersion=${VERSION} -X main.GitCommit=${COMMIT}" \
    -a -installsuffix cgo \
    -o /bin/exporter \
    .


FROM --platform=${TARGETARCH} gcr.io/distroless/static-debian12:latest

LABEL org.opencontainers.image.source=https://github.com/DazWilkin/updown-exporter

COPY --from=build /bin/exporter /

ENTRYPOINT ["/exporter"]
CMD ["--entrypoint=0.0.0.0:8080","--path=/metrics"]
