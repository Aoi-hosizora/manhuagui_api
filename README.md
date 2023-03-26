# manhuagui-backend

+ An unofficial backend for manhuagui (https://www.manhuagui.com/) written in golang/gin.

### Document

+ See http://localhost:10018/v1/swagger/index.html or https://manhuaguibackend.docs.apiary.io/

### Endpoints

+ `GET /v1/manga`
+ `GET /v1/manga/:mid`
+ `GET /v1/manga/:mid/:cid`
+ `GET /v1/list/serial`
+ `GET /v1/list/finish`
+ `GET /v1/list/latest`
+ `GET /v1/list/homepage`
+ `GET /v1/list/updated`
+ `GET /v1/category/genre`
+ `GET /v1/category/zone`
+ `GET /v1/category/age`
+ `GET /v1/category/genre/:genre`
+ `GET /v1/search/:keyword`
+ `GET /v1/author`
+ `GET /v1/author/:aid`
+ `GET /v1/author/:aid/manga`
+ `GET /v1/rank/day`
+ `GET /v1/rank/week`
+ `GET /v1/rank/month`
+ `GET /v1/rank/total`
+ `GET /v1/comment/manga/:mid`
+ `POST /v1/user/login`
+ `POST /v1/user/check_login`
+ `GET /v1/user/info`
+ `GET /v1/user/manga/:mid/:cid`
+ `GET /v1/shelf`
+ `GET /v1/shelf/:mid`
+ `POST /v1/shelf/:mid`
+ `DELETE /v1/shelf/:mid`

### Reference

+ [austinh115/lz-string-go](https://github.com/austinh115/lz-string-go)
