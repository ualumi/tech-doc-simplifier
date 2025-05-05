# Используем официальный образ Go
FROM golang:1.24-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Устанавливаем зависимости, включая git для go get
RUN apk add --no-cache git

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем бинарник
RUN go build -o text-service main.go

# Указываем порт (если нужно, например, для healthcheck; можно убрать)
# EXPOSE 8080

# Команда запуска сервиса
CMD ["./text-service"]
