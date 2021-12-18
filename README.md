# Hotline
DNS/HTTP request logging app

![Dank meme](/meme.jpg)

## Config:

```yml
---
server:
  callback:
    http:
      domain: "mydomain.xyz"
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
    username: "asdf"
    password: "zxcv"
    dbname: "hotline"
client:
  server_url: ""
```