FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY . /app/dota_position_bot/

WORKDIR /app/dota_position_bot/

RUN go mod download
#
RUN go build -o dota_position_bot cmd/main.go

# Финальный образ
FROM alpine:latest

WORKDIR /app/dota_position_bot/

COPY --from=builder /app/dota_position_bot/ /app/dota_position_bot/
ENV TELEGRAM_BOT_TOKEN="7027973131:AAGqw6BZLqa5As5UlPIbRcvnqr0Eehww-ho"
ENV DSN="postgresql://postgres:123123@postgres:5432/postgres?sslmode=disable"
ENV STORAGE_TYPE="postgres"

# Устанавливаем точку входа
ENTRYPOINT ["./dota_position_bot"]