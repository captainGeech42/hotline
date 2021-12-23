# Hotline
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
4. Copy `nginx/nginx_example.conf` to `nginx/nginx.conf` and edit accordingly
4. `docker compose build`
5. `docker compose run`

If you already have something running on port 53, you'll need to stop that service. A common example of this is `systemd-resolved`. To permanently stop `systemd-resolved`, do the following:

```
$ sudo systemctl stop systemd-resolved
$ sudo systemctl disable systemd-resolved
$ sudo mv /etc/resolv.conf /etc/resolv.conf.bak
$ echo -e "nameserver 8.8.8.8\nnameserver 8.8.4.4" | sudo tee /etc/resolv.conf
```

## Testing

To spin up a database without using the `docker-compose.yml` (note that this doesn't use the persistent DB data volume):

```
$ docker run --rm -d -p 3306:3306 --env-file .env mariadb:10.7
```