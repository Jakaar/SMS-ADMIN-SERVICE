<p align="center">
    <a href="https://www.facebook.com/Dr.jawhaa/" target="_blank">
        <img src="https://avatars.githubusercontent.com/u/54393995?v=4" width="100">
    </a>
</p>
<p align="center">
    Developed by Jakaar
</p>

# SMS APP Admin Service v0.0.1-ийн тухай.
`1.`Тус SMS App-ийн Admin service нь GoLang хэл дээр бичэгдсэн. [Go Үзэх](https://go.dev/). <br>
`2.`Мөн Gin gonic Framework ашиглаж хийгдсэн [Gin Gonic Үзэх](https://gin-gonic.com/).



## SMS App Admin Service ийг Development орчинд асаах.
### 1.  Requirement

    1. Go 1.19.2 Xамгийн багадаа.
    2. Postgre SQL

Татах холбоосууд.
<p align="center" >
    <a href="https://go.dev/">
        <img src="https://go.dev/images/go-logo-white.svg" alt="GoLang" width="80">
    </a>
</p>

### 2.  Installation
   ```bash
# 1-р Алхам
$ go get
# 1.1-р Алхам
$ cp .env.dev .env

# 2-р Алхам localhost:8080 
$ go run ./main.go

# Production болсон үед
$ go build -tags=jsoniter main.go
$ ./main
```
[localhost:8080](http://localhost:8080) дээр асна.

[//]: # (#### Username: `admin@admin.com `)

[//]: # (#### Password: `admin123 `)
