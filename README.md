# drone-azure-storage

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-azure-storage/status.svg)](http://cloud.drone.io/drone-plugins/drone-azure-storage)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/azure-storage.svg)](https://microbadger.com/images/plugins/azure-storage "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-azure-storage?status.svg)](http://godoc.org/github.com/drone-plugins/drone-azure-storage)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-azure-storage)](https://goreportcard.com/report/github.com/drone-plugins/drone-azure-storage)
[![](https://images.microbadger.com/badges/image/plugins/azure-storage.svg)](https://microbadger.com/images/plugins/azure-storage "Get your own image badge on microbadger.com")

Drone plugin to publish files and artifacts to Azure Storage. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-azure-storage <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "account_key": "123889asd89u8hsfdjh98128hh",
        "storage_account": "my-storage-account",
        "container": "my-container",
        "source": "folder/to/upload"
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-azure-storage <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "account_key": "123889asd89u8hsfdjh98128hh",
        "storage_account": "my-storage-account",
        "container": "my-container",
        "source": "folder/to/upload"
    }
}
EOF
```
