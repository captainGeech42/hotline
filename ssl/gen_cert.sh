#!/bin/bash

if [[ $# -ne 2 ]]; then
    echo "usage: gen_cert.sh [callback domain] [path to acme challenge response directory]" 2>&1
    exit 2
fi

# make sure docker is available
command -v docker >/dev/null
if [[ $? -ne 0 ]]; then
    echo 'failed to find docker in $PATH' 2>&1
    exit 1
fi

# generate the certificates
#docker run --rm -it -v "$2:/acme-share" certbot-hotline certonly \
docker run --rm -it -v "$2:/acme-share" captaingeech/certbot-dns-hotline:latest certonly \
    --authenticator dns-hotline \
    --dns-hotline-path /acme-share \
    --server https://acme-v02.api.letsencrypt.org/directory \
    --agree-tos \
    --rsa-key-size 4096 \
    -d "$1" \
    -d "*.$1" \
    -v