FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy -compat=1.23 && go mod download

COPY . .

RUN go build -o /app/bin/consumer ./cmd/consumer
RUN go build -o /app/bin/publisher ./cmd/publisher

FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache bash sed wget

COPY --from=builder /app/bin ./bin

RUN chmod +x ./bin/*

