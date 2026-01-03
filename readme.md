# TaskTracker

TaskTracker — распределённый reliable трекер задач, реализованный на Go с использованием микросервисной архитектуры. Проект позволяет регистрировать и авторизовывать пользователей, добавлять задачи и управлять ими, оставлять лайки и комментарии к задачам, а также собирать статистику по просмотрам и лайкам.

Проект содержит 3 отдельных сервиса:

- **userservice** — регистрация и авторизация пользователями  
- **taskservice** — бизнес-логика задач
- **statisticsservice** — сбор и предоставление статистических данных (просмотры, лайки)

---

## Стек использованных технологий

* Go как основной язык проекта
* Docker + Docker Compose для удобного и быстрого развертывания
* grpc и protobuf для взаимодействия между сервисами
* Postgres как основная база данных для хранения информации о пользователях и задачах
* Clickhouse как база данных для сервиса статистики
* Kafka для надёжной передачи данных для statisticsservice из userservice
* REST API, C4 и ER-диаграммы для документации проекта

---

## Структура проекта

```text
├── doc/                       # Документация проекта (API, ER и C4-диаграммы)
├── protos/                    # gRPC/Protobuf файлы
├── userservice/               # Сервис для работы с пользователями
├── taskservice/               # Сервис для работы с задачами задачами
├── statisticsservice/         # Сервис статистики
├── docker-compose.yml         
├── readme.md
└── .gitignore
```

## Установка и запуск 

```bash
$ git clone https://github.com/cutehater/tasktracker.git & cd tasktracker
$ docker compose up --build
```

После этого все сервисы будут подняты и готовы к работе. Сервер будет запущен локально на порту 8008.

## Примеры запросов к API

### Создание пользователя

```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{
    "login": "testuser",
    "password": "password123"
  }'
```

### Авторизация пользователя (получение cookie)
```bash
curl -i -X GET "http://localhost:8080/user/login?login=testuser&password=password123"
```

В ответе будет заголовок `Set-Cookie: JSESSIONID=....`
Этот cookie необходимо передавать в последующих запросах.

### Создание задачи (с авторизацией)
```bash
curl -X POST http://localhost:8080/task \
  -H "Content-Type: application/json" \
  -H "Cookie: JSESSIONID=YOUR_SESSION_ID" \
  -d '{
    "body": "Task body",
    "status": 1
  }'
```

Полное описание API доступно в файле `doc/api.yml`.
