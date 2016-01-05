# Docker image for the Drone Azure Storage plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-azure-storage
#     make deps build docker

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates python py-pip build-base python-dev libffi-dev openssl-dev && \
  rm -rf /var/cache/apk/*
RUN pip install azure-common cryptography azure-servicemanagement-legacy azure-storage blobxfer

ADD drone-azure-storage /bin/
ENTRYPOINT ["/bin/drone-azure-storage"]
