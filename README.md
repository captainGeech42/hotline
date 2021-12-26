# Hotline

[![Go Report Card](https://goreportcard.com/badge/github.com/captainGeech42/hotline)](https://goreportcard.com/report/github.com/captainGeech42/hotline) [![Build Image](https://github.com/captainGeech42/hotline/workflows/Build/badge.svg)](https://github.com/captainGeech42/hotline/actions?query=workflow%3A%22Build%22) [![Docker Hub Publish](https://github.com/captainGeech42/hotline/workflows/Docker%20Hub%20Publish/badge.svg)](https://github.com/captainGeech42/hotline/actions?query=workflow%3A%22Docker+Hub+Publish%22) [![Docker Hub Image](https://img.shields.io/docker/v/captaingeech/hotline?color=blue)](https://hub.docker.com/repository/docker/captaingeech/hotline/general)

DNS/HTTP request logging app

![Dank meme](/meme.jpg)

## Setup Instructions

Hotline is designed to be deployable on a single system with one public IP address, but is granular enough to be deployed across multiple systems/IP addresses. For multiple systems though, you'll need to modify the provided Docker/Docker Compose setup.

### DNS Setup

Hotline can use one or two different domains (if you use one domain, you'll want to configure a subdomain or two).

In this example, we'll use `hotlinewebapp.xyz` as the front-end web app domain, and `hotlinecallback.net` for the callback domain.

1. Configure an A record for `hotlinewebapp.xyz` to point to your Hotline server
2. Configure a NS record for `hotlinecallback.net` that points to `hotlinewebapp.xyz.`

This enables users of your Hotline server to access it through public DNS servers, but DNS callbacks generated for your Hotline server to properly resolve back to your Hotline DNS callback server.

### Server Setup

1. Install Docker + Docker Compose
2. Copy `.env_sample` to `.env` and edit accordingly
3. Copy `hotline_sample.yml` to `hotline.yml` and edit accordingly
4. Copy `nginx/nginx_sample.conf` to `nginx/nginx.conf` and edit accordingly
5. Generate the React SPA production build: `./build_spa.sh`
6. Build the Hotline server image: `docker-compose build`
7. Start the Hotline server + other components: `docker-compose up`

If you already have something running on port 53, you'll need to stop that service. A common example of this is `systemd-resolved`. To permanently stop `systemd-resolved`, do the following:

```
$ sudo systemctl stop systemd-resolved
$ sudo systemctl disable systemd-resolved
$ sudo mv /etc/resolv.conf /etc/resolv.conf.bak
$ echo -e "nameserver 8.8.8.8\nnameserver 8.8.4.4" | sudo tee /etc/resolv.conf
```

If you'd like to setup SSL certificates, please read [these docs](/ssl/setup.md).

### Client Setup

Now that you have a Hotline server running, you can setup a client. First, install Hotline:

```
$ go install github.com/captainGeech42/hotline@latest
```

Then, setup your config in `~/.hotline.yml`:

```yml
---
client:
  server_url: "http://hotlinewebapp.xyz"
```

Now you are ready to start using Hotline!

```
$ hotline client
2021/12/23 08:27:57 Hotline is now active using your new callback: 9g7yx03b2nvy5hpnpo48
2021/12/23 08:27:57 Start making requests!

        $ curl http://9g7yx03b2nvy5hpnpo48.hotlinecallback.net

        $ dig +short TXT 9g7yx03b2nvy5hpnpo48.hotlinecallback.net

===========================================================================

```

## Usage

### Client Usage

```
$ hotline client -h
Run the Hotline client

Usage:
  hotline client [flags]

Flags:
  -h, --help              help for client
  -n, --name string       Existing callback name to use (leave blank to generate a new one)
  -a, --show-historical   Show all previous requests

Global Flags:
  -c, --config string   Path to config file (ignored if $HOTLINE_CONFIG_PATH is set) (default "$HOME/.hotline.yml")
```

If you run `hotline client` without any arguments, a new callback will be generated, and any requests to your callback (DNS or HTTP) will be streamed to your client.

If you specify the name of an existing callback with `-n`/`--name`, it will be used instead of a randomly generated one. However, if that callback doesn't exist, a new, randomly generated one will be used, rather than the one you specified being created and used.

If you use an existing callback and want to see all of the previous requests, set the `-a`/`--show-historical` flag.

### Server Usage

```
$ hotline server -h
Run the Hotline server (set $HOTLINE_APP to configure which server to run)

Usage:
  hotline server [flags]

Flags:
  -h, --help   help for server

Global Flags:
  -c, --config string   Path to config file (ignored if $HOTLINE_CONFIG_PATH is set) (default "$HOME/.hotline.yml")
```

The main configuration item for running a server (besides the `hotline.yml` file) is setting the `$HOTLINE_APP` environment variable to one of the following options:

* `web`: To run the API backend and serve the React SPA frontend
* `http`: To run the HTTP callback service
* `dns`: To run the DNS callback service

The official Hotline Docker image doesn't set this variable, so you'll need to set it when running a Hotline server. The provided `docker-compose.yml` configures this for each container ([example](https://github.com/captainGeech42/hotline/blob/main/docker-compose.yml#L30)).
