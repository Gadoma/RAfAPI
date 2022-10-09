#!/usr/bin/env bash
EXIT_STATUS=0

echo ">>> Starting linter container"
docker compose -f docker-compose.dist.yml --profile ci up -d linter --force-recreate

echo ">>> Linting 'common'"
docker exec \
-w /app/common \
linter \
golangci-lint run \
|| EXIT_STATUS=$?

echo ">>> Linting 'affirmation'"
docker exec \
-w /app/affirmation \
linter \
golangci-lint run \
|| EXIT_STATUS=$?

echo ">>> Linting 'category'"
docker exec \
-w /app/category \
linter \
golangci-lint run \
|| EXIT_STATUS=$?

echo ">>> Linting 'randomAffirmation'"
docker exec \
-w /app/randomAffirmation \
linter \
golangci-lint run \
|| EXIT_STATUS=$?

echo ">>> Stopping linter container"
docker stop linter
docker rm linter

if [ "$EXIT_STATUS" -eq "0" ]; then
echo ">>> OK"
else
echo ">>> ERROR"
fi

exit $EXIT_STATUS
