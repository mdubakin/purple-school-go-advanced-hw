# Random API

Запускаем в 1 терминале:

```sh
go run cmd/main.go
```

Делаем запросы во 2 терминале:

```sh
❯ curl localhost:8888/random/int
1
❯ curl localhost:8888/random/int
5
❯ curl localhost:8888/random/int
0
❯ curl localhost:8888/random/int
1
❯ curl localhost:8888/random/int
4
```
