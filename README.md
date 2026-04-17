# GORM + Gin: базовое CRUD API

Еще раз прочитай перед началом работы:

- GORM: подключение к базе (PostgreSQL): https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
- GORM: модели: https://gorm.io/docs/models.html
- GORM: миграции (AutoMigrate): https://gorm.io/docs/migration.html#Auto-Migration
- GORM: CRUD: https://gorm.io/docs/create.html, https://gorm.io/docs/query.html, https://gorm.io/docs/update.html, https://gorm.io/docs/delete.html
- Gin: быстрый старт: https://gin-gonic.com/en/docs/quickstart/
- Gin: работа с JSON и привязкой тела запроса: https://gin-gonic.com/en/docs/examples/binding-and-validation/

## Релиз 0 — Установка зависимостей и базовый сервер

Установи GORM, драйвер для PostgreSQL и Gin согласно документации.

- Инициализируй проект (`go mod init ...`).
- Установи `gorm.io/gorm`, `gorm.io/driver/postgres`, `github.com/gin-gonic/gin`.
- Создай минимальный сервер на Gin (`gin.Default()` + `router.Run()`), добавь тестовый роут `GET /ping` → `{ "message": "pong" }`.

Проверка:

- Сервер запускается (`go run .`) и по адресу `GET http://localhost:8080/ping` возвращает JSON.

## Релиз 1 — Модель `Student` на `gorm.Model`

Создай структуру `Student`, сразу встроив в неё `gorm.Model`. Добавь 1–2 простых поля (например, `Name`, `Age`) и JSON‑теги, чтобы имена в ответе были в принятом REST API стиле.

Документация:

- Модели: https://gorm.io/docs/models.html
- gorm.Model: https://gorm.io/docs/models.html#gorm.Model

Проверка:

- Проект компилируется, в структуре присутствуют поля из `gorm.Model`.

## Релиз 2 — Подключение к PostgreSQL и AutoMigrate

Настрой подключение к PostgreSQL и автоматически создавай таблицы при старте сервера.

- Используй `gorm.Open(postgres.Open(dsn), &gorm.Config{})`.
- DSN: хост, порт, пользователь, пароль, имя БД, `sslmode=disable`.
- После успешного подключения вызови `AutoMigrate(&Student{})`.

Документация:

- Подключение (PostgreSQL): https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
- AutoMigrate: https://gorm.io/docs/migration.html#Auto-Migration

Проверка:

- При старте сервера таблица для `students` создаётся (можно дополнительно убедиться в pgAdmin).

## Релиз 3 — Модель `Group` и миграции

Создай вторую модель `Group` (минимум: `ID`, `Name`). Нужно также встроить `gorm.Model`. Добавь её в автоматическую миграцию.

Документация:

- Модели: https://gorm.io/docs/models.html
- AutoMigrate: https://gorm.io/docs/migration.html#Auto-Migration

Проверка:

- После рестарта сервера появляется таблица `groups`.

## Релиз 4 — REST эндпоинты для `Student`

Реализуй CRUD эндпоинты. Возвращай корректные коды ответа и JSON.

- POST `/students`

  - Тело запроса: `{ "name": string, "age": number }`
  - Действие: создать студента.
  - Ответ: созданный объект студента, код `201 Created`.
  - Ошибки: `400 Bad Request` (валидация), `500 Internal Server Error` (БД).

- GET `/students/:id`

  - Действие: получить студента по `id`.
  - Ответ: объект студента, код `200 OK`.
  - Ошибки: `404 Not Found` если не найден.

- PATCH `/students/:id`

  - Тело запроса: например `{ "name": string }` (частичное обновление).
  - Действие: обновить указанные поля.
  - Ответ: обновлённый студент, код `200 OK`.
  - Ошибки: `400 Bad Request`, `404 Not Found`.

- DELETE `/students/:id`

  - Действие: удалить студента.
  - Ответ: `{ "message": "deleted" }`, код `200 OK`.
  - Ошибки: `404 Not Found`.

- GET `/students`

  - Действие: вернуть список студентов.
  - Ответ: `{ "students": [...] }`, код `200 OK`.

Подсказки:

- Создание: https://gorm.io/docs/create.html
- Чтение: https://gorm.io/docs/query.html
- Обновление: https://gorm.io/docs/update.html
- Удаление: https://gorm.io/docs/delete.html
- Gin: привязка и валидация JSON: https://gin-gonic.com/en/docs/examples/binding-and-validation/

## Релиз 5 — REST эндпоинты для `Group`

Сделай аналогичные эндпоинты для группы:

- POST `/groups` — создать группу.
- GET `/groups/:id` — получить группу по `id`.
- PATCH `/groups/:id` — частично обновить.
- DELETE `/groups/:id` — удалить.
- GET `/groups` — список групп.

На этом этапе связи между `Student` и `Group` не настраиваем — сфокусируйся на CRUD и корректных ответах API.

## Закрываем задачу

- Протестируй все эндпоинты в Postman, исправь найденные ошибки.
- При необходимости загляни в pgAdmin, чтобы убедиться, что записи действительно создаются/обновляются/удаляются.
- Выгрузи работу на гитхаб и отправь работу на проверку преподавателю.
