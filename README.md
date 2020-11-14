# manhuagui-backend

+ An unofficial backend for manhuagui (https://www.manhuagui.com/) written in golang/gin

### Endpoints

+ `GET /v1/manga`
+ `GET /v1/manga/:mid`
+ `GET /v1/manga/:mid/:cid`
+ `GET /v1/list/serial`
+ `GET /v1/list/finish`
+ `GET /v1/list/latest`
+ `GET /v1/list/updated`
+ `GET /v1/category/genre`
+ `GET /v1/category/zone`
+ `GET /v1/category/age`
+ `GET /v1/category/genre/:genre`
+ `GET /v1/search/:keyword`
+ `GET /v1/author`
+ `GET /v1/author/:aid`
+ `GET /v1/author/:aid/manga`

### Document

+ See http://localhost:10018/v1/swagger/index.html or https://manhuaguibackend.docs.apiary.io/

### Reference

+ [austinh115/lz-string-go](https://github.com/austinh115/lz-string-go)
