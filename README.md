# Общая информация

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

#### dev

```bash
go run . # В .env вставьте свой DATABASE_URL
```

#### prod

```bash
docker compose up --build # python3.10 и выше
```

## Дерево проекта

```text
├─ config
|   └─ Проверка наличия env
|
├─ internal
|   ├─ constants
|   |   └─ Общие константы (urls, время жизни токенов)
|   ├─ db
|   |   └─ Подключение к БД
|   ├─ server
|   |   └─ Подключение к БД
|   |       ├─ [endpoint]
|   |       |   └─ Роутеры и методы конечных точек
|   |       └─ router.go
|   |           └─ Главный роутер
|   └─ types
|       └─ Типы
|
└─ init.sql
    └─ Скрипт для БД
```

# Конечные точки

1. [Healthcheck](#healthcheck)
2. [POST /auth/login](#authlogin)
3. [POST /auth/refresh_token](#authrefresh_token)

## Эндпоинты

### /healthcheck

Пример **GET** запроса:

```text
curl --location 'http://localhost:5000/healthcheck'
```

Пример ответа:

```text
ok
```

### /auth

#### /auth/login

Пример **POST** запроса:

```text
curl --location 'http://localhost:5000/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "guid": "57979158-bc47-490c-87fb-183d9b7a99d4",
    "ip": "188.243.129.142"
}'
```

Пример ответа:

```text
{
    "access_token": "---token---",
    "refresh_token": "---string---"
}
```

#### /auth/refresh_token

Пример **POST** запроса:

```text
curl --location 'http://localhost:5000/auth/refresh_token' \
--header 'Content-Type: application/json' \
--data '{
    "access_token": "---token---",
    "refresh_token": "---string---",
    "ip": "188.243.129.142"
}'
```

Пример ответа:

```text
{
    "access_token": "---token---",
    "refresh_token": "---string---"
}
```
