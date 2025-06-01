# Мини-сервис “Цитатник”

## инструкцией запуска

Команды исполняются из основной директории

### Компиляция в исполняемый файл

```bash
go build ./cmd/main/citatnik.go

./citatnik
```

### Компиляция в Docker контейнер

```bash
docker build --no-cache -t citatnik .
```

## Проверочные команды (curl):
```bash
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```
```bash
curl http://localhost:8080/quotes
``` 
```bash
curl http://localhost:8080/quotes/random
```
```bash
curl http://localhost:8080/quotes?author=Confucius
```
```bash 
curl -X DELETE http://localhost:8080/quotes/1
```