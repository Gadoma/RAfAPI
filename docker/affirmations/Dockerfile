# syntax=docker/dockerfile:1

## Build stage
FROM golang:1.19 AS builder

LABEL maintainer="Piotr Gadziński <dev@gadoma.com>"

WORKDIR /app

COPY . .

WORKDIR /app/internal/affirmation

RUN dpkg --add-architecture amd64 \
        && apt update \
        && apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="x86_64-linux-gnu-gcc" go build -o /app/bin/affirmation main.go

## Deploy stage
FROM --platform=linux/amd64 debian:bullseye-slim AS runner

WORKDIR /app

COPY --from=builder /app/bin/affirmation /app/affirmation

ENTRYPOINT ["/app/affirmation"]
