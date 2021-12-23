# Development Guide

Hotline has five main components:

* CLI client ([code](/internal/client))
* React SPA web app ([code](/internal/web/frontend/spa))
* REST API backend ([code](/internal/web/backend))
* HTTP callback service ([code](/internal/callback/http))
* DNS callback service ([code](/internal/callback/dns))

![System architecture](/arch.png)

## Testing

To spin up a database without using the `docker-compose.yml` (note that this doesn't use the persistent DB data volume):

```
$ docker run --rm -d -p 3306:3306 --env-file .env mariadb:10.7
```

Make sure you set your `hotline.yml` to point to `localhost` for the database, and then you'll be able to connect to it.

To test the DNS callback service, you can use dig with `@localhost` appended to the args, which will bypass the public nameserver configuration for your callback domain.

To test the HTTP callback service, you can set the domain in `/etc/hosts` or set a `Host` header like this (using HTTPie):

```
$ http :/test Host:9g7yx03b2nvy5hpnpo48.hotlinecallback.net
```