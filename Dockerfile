# Builder stage
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/app

# Final stage
FROM alpine:latest

# Копируем бинарник
COPY --from=builder /server /server

# Копируем .env (если нужен)
COPY --from=builder /app/.env /app/.env

# Копируем миграции
COPY --from=builder /app/migrations /app/migrations

# Рабочая директория
WORKDIR /app

EXPOSE 8080


# Запуск приложения
CMD ["/server"]