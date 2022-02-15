# REST API boilerplate with smtp
> REST API boilerplate with smtp routes
## Установка
#### 1. Клонируем репозиторий
```bash
$ git clone git@gitlab.com:OrlangurDux/rest-api-boilerplate-with-smtp.git  
```
#### 2. Устанавливаем вендорные зависимости
```bash
$ go mod vendor
```
## Разработка

### Запуск тестового сервера

```bash
$ go run main.go routes.go
```
Запускаем сервисы, результаты в 
* 🌏 **API Сервер** запущен на `http://localhost:9099`
* ⚙️**Swagger UI** запущен на `http://localhost:9099/swagger/`

## Упаковка и Развертывание
#### 1. Сборка и запуск без Docker
```bash
$ go build 
```
#### 2. Запуск тестов
```bash
$ cd tests
$ go test
```
#### 3. Запуск с Docker
```bash
$ docker build -t api-server .
$ docker run -t -i -p 9099:9099 api-server
```
#### 4. Запуск с Docker-Compose
```
$ docker-compose up
```

>Примечание в случае запуска не через Docker-Compose необходимо поднять отдельно MongoDB и скорректировать файл .env

---

## Окружение
Для редактирования переменных окружения, создайте файл с именем `.env` и скопируйте содержимое из `.env.default` чтобы начать.

| Var Name  | Type  | Default | Description  |
|---|---|---|---|
|  PORT | number  | `9099` | Порт на котором запускается API сервер |
|  SMTP_HOST | string  | `smtp.mail.loc` | SMTP шлюз |
|  SMTP_LOGIN | string  | `sender@mail.loc` | Логин |
|  SMTP_PASSWORD | string  | `password` | Пароль |
|  SMTP_PORT | number  | `587` | Порт для подключения 587, 25 |
|  SMTP_MESSAGE | string  | `Message mail` | Тело письма поддерживает html и использует метки из запроса #переменная_в_запросе# |
|  SMTP_SUBJECT | string  | `Subject` | Тема письма |
|  SMTP_TO | string  | `recepient@mail.loc` | Получатель |
