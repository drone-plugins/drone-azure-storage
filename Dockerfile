# Docker image for the Drone Azure Storage plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-azure-storage
#     make deps build docker

FROM alpine:3.3

RUN apk update && \
  apk add \
    ca-certificates \
    python \
    py-pip \
    build-base \
    python-dev \
    libffi-dev \
    openssl-dev && \
  pip install --upgrade \
    pip && \
  pip install \
    azure-common>=1.0.0 \
    requests>=2.9.1 \
    cryptography>=1.2.2 \
    azure-servicemanagement-legacy>=0.20.1 \
    azure-storage>=0.20.3 \
    blobxfer==0.9.9.11 && \
  rm -rf /var/cache/apk/*

ADD drone-azure-storage /bin/
ENTRYPOINT ["/bin/drone-azure-storage"]
