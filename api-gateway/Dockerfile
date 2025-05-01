# Базовый образ с Go
FROM golang:1.24-alpine

# Создаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum, затем устанавливаем зависимости (кеш эффективно)
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN go build -o api-gateway .

# Указываем порт, на котором работает сервис
EXPOSE 8080

# Указываем команду запуска
CMD ["./api-gateway"]
