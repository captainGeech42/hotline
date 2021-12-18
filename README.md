# Hotline
DNS/HTTP request logging app

![Dank meme](/meme.jpg)

## Config:

For a client, you can omit the `server` block. The below sample config shows all possible values

```yml
---
server:
  callback:
    domain: "mydomain.xyz"
    http:
      port: 8080
      default_response: "research by @captainGeech using hotline"
    dns:
      port: 53
      default_A_response: 1.2.3.4
      default_TXT_response: "research by @captainGeech using hotline"
  app:
    port: 8080
  db:
    host: "localhost"
    port: 3306
    username: "dbuser"
    password: "put_a_secure_pass_here"
    dbname: "hotline"
client:
  server_url: "http://otherdomain.abc"
```