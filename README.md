# golang-webapi-practice

```console
$ docker-compose up -d
$ go run server.go database.go
```

```console
$ curl -Ssv -X POST "localhost:1323/users" -H 'Content-Type: application/json' -d '{"name":"hoge"}' | jq
$ curl -Ssv "localhost:1323/users" | jq
```