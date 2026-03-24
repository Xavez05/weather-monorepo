# Stage 1: build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar ambos módulos del monorepo
COPY apiclient/ ./apiclient/
COPY weather-app/ ./weather-app/
COPY go.work .

# Build del binario
WORKDIR /app/weather-app
RUN go build -o ../bin/weather-app ./cmd/main.go

# Stage 2: imagen final liviana
FROM alpine:latest

WORKDIR /app

# Copiar binario compilado
COPY --from=builder /app/bin/weather-app .

# Copiar templates HTML
COPY --from=builder /app/weather-app/internal/templates/ ./internal/templates/

EXPOSE 8080

CMD ["./weather-app"]