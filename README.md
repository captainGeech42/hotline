# Hotline
DNS/HTTP request logging app

![Dank meme](/meme.jpg)

## Config

For a client, you can omit the `server` block. The below sample config shows all possible values. Please note that the default Dockerfile only exposes port 8080/tcp and 53/[tcp,udp].

```yml
---
server:
  callback:
    domain: "mydomain.xyz"
    http:
      port: 8080
      default_response: "research by @captainGeech42 using hotline"
    dns:
      port: 53
      default_A_response: 1.2.3.4 # this needs to be the public IP for the hotline HTTP callback server
      default_TXT_response: "research by @captainGeech42 using hotline"
  app:
    port: 8080
  db:
    host: "db"
    port: 3306
    username: "dbuser"
    password: "put_a_secure_pass_here"
    dbname: "hotline"
client:
  server_url: "http://otherdomain.abc"
```

## Testing

To spin up a database without using the `docker-compose.yml` (note that this doesn't use the persistent DB data volume):

```
$ docker run --rm -d -p 3306:3306 --env-file .env mariadb:10.7
```