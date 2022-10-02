#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f "$0")")

docker run --rm -v $BASE_DIR/../:/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run -v
