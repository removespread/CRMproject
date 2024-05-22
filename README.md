# CRMproject

# Базовые требования к проекту:
- Завести приватный репозиторий, добавить Qwent71 и pers0na2dev в репо. 
- Монолитный сервис на Golang 1.22.
- Сервер - net/http, роутинг - net/http.
- Использовать fx для инжекта зависимостей (https://github.com/uber-go/fx)
- Запросы в базу данных через gorm (https://gorm.io/index.html)
- База данных - PostgreSQL запущенная в Docker-compose.
- Обязательно слоистая архитектура с разбиением на репозитории, сервисы и хендлеры.
- Миграции через Goose (https://github.com/pressly/goose) 
- Авторизация пользователя через JWT
- Конфиг приложения через Viper (https://github.com/spf13/viper)
# Техническое задание

4 базовые сущности

Account:
- ID - primary key, int, autoincrement
- Email - string
- Password - string

Contact:
- ID - primary key, int, autoincrement
- Name - string
- Phone - string
- Description - string

Partner:
- ID - primary key, int, autoincrement
- Name - string
- Description - string
- []Contacts - string, контакты у контрагента

Bid:
- ID - primary key, int, autoincrement
- Description - string
- Amount - int
- CreatedAt - timestamp

Хендлеры:
POST /auth/signup - регистрация в сервисе
POST /auth/signin - авторизация в сервисе

GET /contact - список всех контактов
GET /contact/:id - получить один контакт
POST /contact - создать контакт
PUT /contact/:id - обновить контакт
DELETE /contact/:id - удалить контакт

GET /partner - список всех партнеров
GET /partner/:id - получить одного партнера
POST /partner - создать партнера
PUT /partner/:id - обновить партнера
DELETE /partner/:id - удалить партнера

GET /bid - список всех заявок
GET /bid/:id - получить одну заявку
POST /bid - создать заявку
PUT /bid/:id - обновить заявку
DELETE /bid/:id - удалить заявку
