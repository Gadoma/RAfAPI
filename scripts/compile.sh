#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f "$0")")

env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" \
go build -o $BASE_DIR/../bin/rafapi $BASE_DIR/../cmd/rafapi/main.go
