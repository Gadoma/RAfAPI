#!/usr/bin/env bash
BASE_DIR=$(dirname "$(readlink -f "$0")")
EXIT_STATUS=0

echo ">>> Building 'common'"
cd $BASE_DIR/../internal/common
go build ./... || EXIT_STATUS=$?

echo ">>> Building 'affirmation'"
cd $BASE_DIR/../internal/affirmation
go build ./... || EXIT_STATUS=$?

echo ">>> Building 'category'"
cd $BASE_DIR/../internal/category
go build ./... || EXIT_STATUS=$?

echo ">>> Building 'randomAffirmation'"
cd $BASE_DIR/../internal/randomAffirmation
go build ./... || EXIT_STATUS=$?

if [ "$EXIT_STATUS" -eq "0" ]; then
echo ">>> OK"
else
echo ">>> ERROR"
fi

exit $EXIT_STATUS
