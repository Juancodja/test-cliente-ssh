FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o sushi .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sushi-ssh .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sushi-ssh /app/sushi-ssh

ENTRYPOINT ["/app/sushi", "-p", "22", "test@192.168.1.10" ]
#ENTRYPOINT ["tail", "-f", "/dev/null"]

