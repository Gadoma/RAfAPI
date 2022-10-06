#!/usr/bin/env bash
BASE_DIR=$(dirname "$(readlink -f "$0")")
EXIT_STATUS=0

echo ">>> Linting 'common'"
docker run \
--rm \
-v $BASE_DIR/../internal/common/:/app/common \
-w /app/common \
golangci/golangci-lint:v1.49.0 \
golangci-lint run || EXIT_STATUS=$?

echo ">>> Linting 'affirmation'"
docker run \
--rm \
-v $BASE_DIR/../internal/common/:/app/common \
-v $BASE_DIR/../internal/affirmation/:/app/affirmation \
-w /app/affirmation \
golangci/golangci-lint:v1.49.0 \
golangci-lint run || EXIT_STATUS=$?

echo ">>> Linting 'category'"
docker run \
--rm \
-v $BASE_DIR/../internal/common/:/app/common \
-v $BASE_DIR/../internal/category/:/app/category \
-w /app/category \
golangci/golangci-lint:v1.49.0 \
golangci-lint run || EXIT_STATUS=$?

echo ">>> Linting 'randomAffirmation'"
docker run \
--rm \
-v $BASE_DIR/../internal/common/:/app/common \
-v $BASE_DIR/../internal/randomAffirmation/:/app/randomAffirmation \
-w /app/randomAffirmation \
golangci/golangci-lint:v1.49.0 \
golangci-lint run || EXIT_STATUS=$?

if [ "$EXIT_STATUS" -eq "0" ]; then
echo ">>> OK"
else
echo ">>> ERROR"
fi

exit $EXIT_STATUS
