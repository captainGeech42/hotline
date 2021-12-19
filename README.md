# Hotline
DNS/HTTP request logging app

![Dank meme](/meme.jpg)

## Setup Instructions

### DNS Setup

tbd

### Server Setup

1. Install Docker + Docker Compose
2. Copy `.env_sample` to `.env` and edit accordingly
3. Copy `hotline_sample.yml` to `hotline.yml` and edit accordingly
4. Copy `nginx/nginx_example.conf` to `nginx/nginx.conf` and edit accordingly
4. `docker compose build`
5. `docker compose run`

## Testing

To spin up a database without using the `docker-compose.yml` (note that this doesn't use the persistent DB data volume):

```
$ docker run --rm -d -p 3306:3306 --env-file .env mariadb:10.7
```