# REST-API приложение "Онлайн-библиотека"🚀
Сервис, хранящий пользователей и их приватные и общедостпуные книги.

Используемые технологии:
+ PostgreSQL (хранилище данных)
+ Docker (запуск сервиса)
+ Gin (веб-фреймворк)
+ pq (драйвер работы с PostgreSQL)
+ logrus (логирование)
+ JWT-токен (аутентификация)

## 🔧Getting Started
Для запуска данного приложения вам потребуется Docker

## Запуск 🛠️
Для запуска выполните в терминале команду make run, после чего сервер будет запущен на localhost на порту 8082. Для остановки сервера нужно прописать команду make stop

## REST Методы
### POST /user-auth/sign-up (Регистрация пользователя)
Пример запроса:
```
{
  "username":"john",
  "password":"mypass",
  "email":"example@yandex.ru
}
```
Пример ответа:
```
200 OK (успешно)
400 BadRequest (некорректное тело запроса)
409 Conflict (username уже существует)
```

### POST /user-auth/sign-in (Вход пользователя в систему)
Пример запроса:
```
{
  "username":"john",
  "password":"mypass"
}
```
Пример ответа:
```
200 OK (успешно)
400 BadRequest (некорректное тело запроса)
401 Unathorizes (пользователь не зарегистрирован)
```

### GET /api/user (Получение информации о пользователе)
Пример ответа:
```
{
  "username":"john",
  "email":"example@yandex.ru
}
```
Коды ответа:
```
200 OK
400 BadRequest
500 InternalServerError
```

### POST /api/user (Обновление информации о пользователе)
```
{
  "username":"john",
  "password":"mypass",
  "email":"example@yandex.ru"
}
```
Коды ответа:
```
200 OK
400 BadRequest
409 Conflict
500 InternalServerError
```

### DELETE /api/user (Удаление пользователя)
Коды ответа:
```
200 OK
400 BadRequest
500 InternalServerError
```

### GET /api/{username}/books (Получение информации о книгах пользователя)
Пример ответа:
```
[
	{
		"title":"book1",
    "author":"author1",
    "issue_year":2024
	},
	{
		"title":"book2",
    "author":"author2",
    "issue_year":2024
	},
]
```
Коды ответа:
```
200 OK
400 BadRequest
404 NotFound
500 InternalServerError
```

### GET /api/books/{bookId} (Скачивание книги)
Коды ответа:
```
200 OK
400 BadRequest
404 NotFound
500 InternalServerError
```

### POST /api/book (Создание новой книги)
Пример запроса:
```
{
  "title":"book1",
  "author":"author1",
  "issue_year":2024,
  "file_name":"b60fd97d-5823-4bc9-9df0-8f951cb1231b.pdf",
  "access":"private"
}
```
Коды ответа:
```
200 OK
400 BadRequest
500 InternalServerError
```

### POST /api/books/upload (Загрузка книги)
Коды ответа:
```
200 OK
400 BadRequest
500 InternalServerError
```
### POST /api/books/{bookId} (Обновление данных книги)
Пример запроса:
```
{
  "title":"book1",
  "author":"author1",
  "issue_year":2024,
  "access":"private"
}
```
Коды ответа:
```
200 OK
400 BadRequest
404 NotFound
500 InternalServerError
```

### DELETE /api/books/{bookId} (Удаление книги)
Коды ответа:
```
200 OK
400 BadRequest
500 InternalServerError
```
