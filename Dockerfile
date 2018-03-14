FROM plugins/base:amd64

LABEL maintainer="Drone.IO Community <drone-dev@googlegroups.com>" \
  org.label-schema.name="Drone Azure Storage" \
  org.label-schema.vendor="Drone.IO Community" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache python py-pip build-base python-dev libffi-dev openssl-dev && \
  pip install azure-common>=1.0.0 requests>=2.9.1 cryptography>=1.2.2 azure-servicemanagement-legacy>=0.20.1 azure-storage>=0.20.3 blobxfer==0.9.9.11

ADD release/linux/amd64/drone-azure-storage /bin/
ENTRYPOINT ["/bin/drone-azure-storage"]
