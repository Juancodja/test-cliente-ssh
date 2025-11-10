FROM golang:1.24 AS builder

WORKDIR /app

# Copiamos solo dependencias primero para mejor cache
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o sushi .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sushi /app/sushi

ENTRYPOINT ["/app/sushi"]

