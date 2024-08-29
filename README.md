# part_authentication_service

Часть сервиса аутентификации

## Задание

Написать часть сервиса аутентификации.

Два REST маршрута:

- Первый маршрут выдает пару Access, Refresh токенов для пользователя с идентификатором (GUID), указанным в параметре запроса.
- Второй маршрут выполняет Refresh операцию на пару Access, Refresh токенов.

### Требования

- Access токен: тип JWT, алгоритм SHA512, хранить в базе строго запрещено.

- Refresh токен: тип произвольный, формат передачи base64, хранится в базе исключительно в виде bcrypt хеша, должен быть защищен от изменения на стороне клиента и попыток повторного использования.

- Access, Refresh токены обоюдно связаны, Refresh операцию для Access токена можно выполнить только тем Refresh токеном который был выдан вместе с ним.

- Payload токенов должен содержать сведения об ip адресе клиента, которому он был выдан. В случае, если ip адрес изменился, при рефреш операции нужно послать email warning на почту юзера (для упрощения можно использовать моковые данные).

## Запуск

### Первый шаг

```bash
cp env.example .env
```

### Второй шаг

#### prod

```bash
docker compose up --build # python3.10 и выше
```
