#!/bin/bash

if [[ $# -ne 2 ]]; then
    echo "usage: gen_cert.sh [callback domain]" 2>&1
    exit 2
fi

# make sure docker is available
command -v docker >/dev/null
if [[ $? -ne 0 ]]; then
    echo 'failed to find docker in $PATH' 2>&1
    exit 1
fi

# pull the latest image
docker pull captaingeech/certbot-dns-hotline:latest

# generate the certificates
docker run --rm -it -v "$(pwd)/dns_chal:/acme-share" -v "$(pwd)/callback_cert:/etc/letsencrypt/live/$1" captaingeech/certbot-dns-hotline:latest certonly \
    --authenticator dns-hotline \
    --dns-hotline-path /acme-share \
    --server https://acme-v02.api.letsencrypt.org/directory \
    --agree-tos \
    --rsa-key-size 4096 \
    -d "$1" \
    -d "*.$1" \
    -v