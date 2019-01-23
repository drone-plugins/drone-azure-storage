# drone-azure-storage

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-azure-storage/status.svg)](http://cloud.drone.io/drone-plugins/drone-azure-storage)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/azure-storage.svg)](https://microbadger.com/images/plugins/azure-storage "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-azure-storage?status.svg)](http://godoc.org/github.com/drone-plugins/drone-azure-storage)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-azure-storage)](https://goreportcard.com/report/github.com/drone-plugins/drone-azure-storage)

Drone plugin to publish files and artifacts to Azure Storage. For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-azure-storage/).

## Build

Build the binary with the following commands:

```
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go test -cover ./...
go build -v -a -tags netgo -o release/linux/amd64/drone-azure-storage
```

## Docker

Build the Docker image with the following commands:

```
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/azure-storage .
```

### Usage

```
docker run --rm \
  -e PLUGIN_ACCOUNT=my-storage-account \
  -e PLUGIN_ACCOUNT_KEY=123889asd89u8hsfdjh98128hh \
  -e PLUGIN_CONTAINER=my-container \
  -e PLUGIN_SOURCE=folder/to/upload \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/azure-storage
```
