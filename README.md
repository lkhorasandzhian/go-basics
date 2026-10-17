# Go Basics

Репозиторий содержит решения 10 практических задач на Go, охватывающих условные операторы, циклы, функции и обработку пользовательского ввода.

Каждое задание расположено в отдельной папке `task-01` — `task-10`.

## Запуск

```bash
docker compose up -d
docker compose exec go go run ./task-01/main.go
```

Для запуска другого задания измените номер папки.

## Остановка

```bash
docker compose down
```

## Технологии

- Go 1.27
- Docker Compose
- VS Code + Dev Containers
