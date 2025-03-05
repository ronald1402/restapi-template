FROM golang:1.23.2-alpine AS builder

RUN apk update && apk add --no-cache git gcc g++ libc-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o restapi ./cmd/service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/restapi .

ENTRYPOINT ["/app/restapi"]
