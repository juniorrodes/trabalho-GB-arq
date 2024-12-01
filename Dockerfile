FROM docker.io/golang:1.23.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o app-bin pkg/main.go

FROM docker.io/alpine:3.14

WORKDIR /app

COPY --from=builder /app/app-bin ./app

CMD ["./app"]
