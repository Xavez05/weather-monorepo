# Stage 1: build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar ambos módulos del monorepo
COPY apiclient/ ./apiclient/
COPY weather-app/ ./weather-app/
COPY go.work .

# Build del binario — ahora main.go está en la raíz de weather-app
WORKDIR /app/weather-app
RUN go build -o ../bin/weather-app .

# Stage 2: imagen final liviana
FROM alpine:latest

WORKDIR /app

# Solo el binario — ya no hay templates
COPY --from=builder /app/bin/weather-app .

EXPOSE 8080

CMD ["./weather-app"]