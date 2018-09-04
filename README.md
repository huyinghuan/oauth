## Docker private registry auth server

## How to use

see `docker-compose.yaml`

create a schame and set mysql config  

## how to registry an account

```
curl -X POST \
  http://127.0.0.1:8000/__registry \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F username=test \
  -F password=123456
```

## Dev

```
cd $GOPATH/src
git clone git@github.com:huyinghuan/oauth.git
go run main.go
```

## How to use register V2 Restful API

Example, if you want to get all repositories

register restful api: https://docs.docker.com/registry/spec/api/#detail

### First step
```
curl -i "http://localhost:5000/v2/_catalog?n=20&last="
```

Result :
```
HTTP/1.1 401 Unauthorized
Content-Type: application/json; charset=utf-8
Docker-Distribution-Api-Version: registry/2.0
Www-Authenticate: Bearer realm="http://172.28.210.141:8000/authorize",service="Docker registry",scope="registry:catalog:*"
X-Content-Type-Options: nosniff
Date: Wed, 08 Aug 2018 06:34:03 GMT
Content-Length: 145

{"errors":[{"code":"UNAUTHORIZED","message":"authentication required","detail":[{"Type":"registry","Class":"","Name":"catalog","Action":"*"}]}]}

```

Here you can get some value from HTTP Header `Www-Authenticate`:

```
realm: http://172.28.210.141:8000/authorize
service: Docker registry
scope: registry:catalog:*
```

and then will use them in next step.


### Second step:

if your account(eg. username: `admin`, password: `123456`) have permission for the `scope` : `registry:catalog:*`

```
fake code:
curl -i -u $username:$password -d "service=$service" -d "scope:$scope" -d "account=$username" $realm 

example code:
curl -i -u admin:123456 -d "service=Docker registry" -d "scope=registry:catalog:*" -d "account=admin" "http://172.28.210.141:8000/authorize"
```

Result:
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 08 Aug 2018 06:43:20 GMT
Content-Length: 1095

{"token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IkVQNVQ6NU1KSTpTUjZEOjNTNFI6RVJGWjpOSUJHOkpWR1c6NlFXVzpPTE9JOlRRTlg6NU1DTzoyQklZIn0.eyJpc3MiOiJBdXRoIFNlcnZpY2UiLCJzdWIiOiJhZG1pbiIsImF1ZCI6IkRvY2tlciByZWdpc3RyeSIsImV4cCI6MTUzMzcxMTUwMCwibmJmIjoxNTMzNzEwNTkwLCJpYXQiOjE1MzM3MTA2MDAsImp0aSI6IjY5NzcwMzkwNzU2NzExMDA1MzciLCJhY2Nlc3MiOlt7InR5cGUiOiJyZWdpc3RyeSIsIm5hbWUiOiJjYXRhbG9nIiwiYWN0aW9ucyI6WyIqIl19XX0.UXmzFsnVGEcKk3xmMHC9eEeAPaQZcxQpdgkLFUab86P7gRhQL7cl9zq4nWg01MK2EmHZCHMr7wu1ql4V0-43lyL6Eg8mV7J8ZrgXy_kDEHWN44RsaEtQOxDmJ6qHJMTI2xm4b2DffqcvFr5IVwSk6qIcwyVGRCKVflvfEwvQuKZGwkOwuaZVqUxVG88Eorff3WZ0ZAw15AkQbZxD1qVpqEILllSjeKZMCB-epMzRAcOynIr4ZiGUqUMUwHF0qI5iQthguTUN8nEN8lJGh7fgH_fEt55NNQ07vANMk21TTxTrQfFR2qoaUs3VsCxhgtI3wqQ5m4ddVqWeKF6cVTJ_ib3TTmn4U0kJJqOBzj1MpTz9ducWjIEO_eJq3p3IfhCqGHe2Qslgmr0HJ1j2oZ3L8N97JVLAYhCNoL82Py7BQmzXJKfZMTkLnokmdYG_WIMfMSNRQfx8NvljtzJtG1OXJ8wsVB5qga66fqod9krPvDEYDhYenkhm3Fq5nKq88EAkpuR0AjSIYQfPazsv6R-hStt7b-TpjKqGcoO3MNt2U0YhCx-8Xu0kBDWU_5vie_w1azJ-AWTxnsQM7_vJLu_0n7JpqLTlo-kMUGuR8a0moloDdcZlqPsdBhIMAgSmjU1HGmHYBe53oeNtZk_QplEYnHEZOSn-voRgadGHYEbFKrE"}⏎
```

you can get a token if http status code is `200`, else  your account is error or don't have permission for the `scope`

### Last step

```
curl -H "Authorization: bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IkVQNVQ6NU1KSTpTUjZEOjNTNFI6RVJGWjpOSUJHOkpWR1c6NlFXVzpPTE9JOlRRTlg6NU1DTzoyQklZIn0.eyJpc3MiOiJBdXRoIFNlcnZpY2UiLCJzdWIiOiJhZG1pbiIsImF1ZCI6IkRvY2tlciByZWdpc3RyeSIsImV4cCI6MTUzMzcxMTUwMCwibmJmIjoxNTMzNzEwNTkwLCJpYXQiOjE1MzM3MTA2MDAsImp0aSI6IjY5NzcwMzkwNzU2NzExMDA1MzciLCJhY2Nlc3MiOlt7InR5cGUiOiJyZWdpc3RyeSIsIm5hbWUiOiJjYXRhbG9nIiwiYWN0aW9ucyI6WyIqIl19XX0.UXmzFsnVGEcKk3xmMHC9eEeAPaQZcxQpdgkLFUab86P7gRhQL7cl9zq4nWg01MK2EmHZCHMr7wu1ql4V0-43lyL6Eg8mV7J8ZrgXy_kDEHWN44RsaEtQOxDmJ6qHJMTI2xm4b2DffqcvFr5IVwSk6qIcwyVGRCKVflvfEwvQuKZGwkOwuaZVqUxVG88Eorff3WZ0ZAw15AkQbZxD1qVpqEILllSjeKZMCB-epMzRAcOynIr4ZiGUqUMUwHF0qI5iQthguTUN8nEN8lJGh7fgH_fEt55NNQ07vANMk21TTxTrQfFR2qoaUs3VsCxhgtI3wqQ5m4ddVqWeKF6cVTJ_ib3TTmn4U0kJJqOBzj1MpTz9ducWjIEO_eJq3p3IfhCqGHe2Qslgmr0HJ1j2oZ3L8N97JVLAYhCNoL82Py7BQmzXJKfZMTkLnokmdYG_WIMfMSNRQfx8NvljtzJtG1OXJ8wsVB5qga66fqod9krPvDEYDhYenkhm3Fq5nKq88EAkpuR0AjSIYQfPazsv6R-hStt7b-TpjKqGcoO3MNt2U0YhCx-8Xu0kBDWU_5vie_w1azJ-AWTxnsQM7_vJLu_0n7JpqLTlo-kMUGuR8a0moloDdcZlqPsdBhIMAgSmjU1HGmHYBe53oeNtZk_QplEYnHEZOSn-voRgadGHYEbFKrE"  "http://localhost:5000/v2/_catalog?n=20&last="

```

you can get a result from the api.



## How to develop yourself Access control

```
git clone git@github.com:huyinghuan/oauth.git
git checkout nodatabase
```

and then edit `controller/authorize.go` functions `verityUsernameAndPassword` and `verityUserScopePermission`

## How create certificate

you must update dir `certs` files.

```
cd certs
openssl req -newkey rsa:4096 -nodes -sha256 -keyout auth.key -x509 -days 365 -out auth.crt
```

## TODOS

- [x] basic auth process
- [x] add mysql manager account
- [ ] add manager ui

## Other

- docker-compose:  https://github.com/huyinghuan/docker-images/tree/master/register

## Thanks

- `https://github.com/cesanta/oauth`
- `http://ylzheng.com/2016/11/29/docker-registry-details/`
- `https://www.cakesolutions.net/teamblogs/docker-registry-api-calls-as-an-authenticated-user`