**Платежный шлюза на Go**

* Go
* PostgreSQL
* ~~Kafka~~
* ~~Prometheus~~
* ~~Grafana~~
* ~~Jaeger~~
* Docker

**Установка:**

```cli
git clone git@github.com:smbya/PayWay.git .
docker-compose up -d
```


**Запросы:**

**TODO:**

- [X] Добавить общий конфиг
- [x] Добавить в readme инструкцию по запуску и запросам
- [x] Добавить логирование на всех основных уровнях (взять slog)
- [X] Добавить миграции (golang-migrate/migrate)
- [x] Сделать нормальную обработку данных, убрать захардкоженые данные
- [x] Проверить экспортированные структуры на необходимость быть публичными. Убрать где это не нужно
- [ ] Использовать интерфейсы только там, где это необходимо. В других местах убрать
- [ ] сделать ближе к прод-реди решению (то до чего руки не дошли)
- [x] merge webserver and controller in file struct
- [x] Перенести /db/ в `internal/repository`,

> * internal/repository/repository.go - только интефрейсы
> * internal/repository/posgresql/implement.go
> * internal/repository/posgresql/db/

- [x] Переиспользование, разделение сервисов. Поправить выходные параметры
- [ ] Создать handler model в контроллере
- [ ] Сделать монорепу


**Команды:**

```bash
# Применить все миграции
make migrate-up

# Откатить последнюю миграцию
make migrate-down

# Создать новую миграцию
make migrate-create name=create_users_table

# Пересоздать все таблицы заново
make migrate-reload

# Сгенерировать sqlc код
make sqlc

# Пересобрать и запустить web сервис
make web-rebuild
```

**API запросы:**

```bash
# Создать платеж
curl -X POST http://localhost:80/payments \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 123,
    "amount": "100.50",
    "currency": "RUB",
    "destination": "wallet123",
    "description": "Оплата заказа"
  }'

# Получить платеж по ID
curl -X GET http://localhost:80/payments/{id}
```


===

**План проекта**

[Task.md](assets/task.md)

![](assets/tg_image_2563737507.png)
