# Snippetbox

Веб-сервис для публикации коротких текстовых сообщений (сниппетов) с регистрацией пользователей и временем жизни записей.

<img width="2548" height="1289" alt="Image" src="https://github.com/user-attachments/assets/ab041c64-3453-4551-8189-5e4117823f40" />

## Возможности

- Создание и просмотр сниппетов со сроком жизни
- Регистрация, вход и выход, сессии в MySQL
- Просмотр информации об аккаунте, смена пароля

Вид для зарегистрированного пользователя 

<img width="2548" height="1289" alt="Image" src="https://github.com/user-attachments/assets/6d706c34-7736-4532-ad7a-d62cf2d9b4bf" />

## Технологии

Окно создания сниппета 

<img width="2548" height="1289" alt="Image" src="https://github.com/user-attachments/assets/40c4faab-d08d-47e9-b713-b36c76e266cf" />

Go 1.26, MySQL 8.4, Docker Compose

## Быстрый старт

Нужны Docker и Docker Compose.

```bash
git clone https://github.com/Hofmann250/snippetbox.git

cd snippetbox

docker compose up --build
```

Открой [https://localhost:4000](https://localhost:4000). Браузер предупредит о самоподписанном сертификате, это нормально.

Остановить: `docker compose down`. Вместе с данными БД: `docker compose down -v`.

## Настройка

| Флаг   | По умолчанию | Описание                             |
| ---------- | ----------------------- | -------------------------------------------- |
| `-addr`  | `:4000`               | Адрес HTTP-сервера               |
| `-dsn`   | см. compose           | Строка подключения к MySQL |
| `-debug` | `false`               | Режим отладки                    |

## Структура проекта

```
cmd/web/    точка входа, обработчики, маршруты
internal/   модели и вспомогательные пакеты
ui/         шаблоны и статика
tls/        tls сертификаты
```

## Тесты

```bash
go test ./...
```

## Благодарности

Проект создан по книге Алекса Эдвардса «Let's Go».

## Лицензия

MIT
