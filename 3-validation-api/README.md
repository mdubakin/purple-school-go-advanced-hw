# Verify API

## Конфигурация

Для работы приложения нужен конфиг, должен находится в `config/config.yaml`. Пример конфигурации есть в `config/config.yaml.example`

## Пример работы

1. Отправляем письмо:

```sh
❯ curl -v -XPOST -d '{"email": "<your_email>"}' localhost:8088/send
verify your email
```

2. Пробуем передать неверный hash:

```sh
❯ curl localhost:8088/verify/random_hash
false
```

3. Передаем верный hash:

```sh
❯ curl localhost:8088/verify/6fe6c42a3e97984a19a62e72d1b02fb3335f0cb2d5048d32652a5b6346a835d3f144e04b4487b381745502815b8dddea21c90713f5e0d14c46ab9b2cc01720bc
hello <your_email>: true
```
