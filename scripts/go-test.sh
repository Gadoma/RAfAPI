#!/usr/bin/env bash
BASE_DIR=$(dirname "$(readlink -f "$0")")
EXIT_STATUS=0

echo ">>> Running tests for 'common'"
cd $BASE_DIR/../internal/common
go test -count=1 ./... || EXIT_STATUS=$?

echo ">>> Running tests for 'affirmation'"
cd $BASE_DIR/../internal/affirmation
go test -count=1 ./...  || EXIT_STATUS=$?

echo ">>> Running tests for 'category'"
cd $BASE_DIR/../internal/category
go test -count=1 ./... || EXIT_STATUS=$?

echo ">>> Running tests for 'randomAffirmation'"
cd $BASE_DIR/../internal/randomAffirmation
go test -count=1 ./... || EXIT_STATUS=$?

if [ "$EXIT_STATUS" -eq "0" ]; then
echo ">>> OK"
else
echo ">>> ERROR"
fi

exit $EXIT_STATUS
