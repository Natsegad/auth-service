# auth-service

В качестве архитектуры приложения используется Чистая архитектура.

Для авторизации пользователя на сервер оправляется POST запрос на /login. Сервер парсит form далее заносит в базу данных postgresql.
Ответом на регистрацию будет json в виде токена пользователя. JWT токен будет работать 1 час. В будущем будет добавлен рефреш токен.

стек: Postgresql, gorm, gin
