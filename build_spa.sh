#!/bin/bash

set -e

# make sure docker is available
command -v docker >/dev/null
if [[ $? -ne 0 ]]; then
    echo 'failed to find docker in $PATH' 2>&1
    exit 1
fi

# build the docker image
docker build -t hotline-spa-build -f Dockerfile.spa .

# build the app
docker run --rm -it -v "$(pwd)/nginx/spa:/app/build" hotline-spa-build npm run build