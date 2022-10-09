#!/usr/bin/env bash
EXIT_STATUS=0

echo ">>> Curling the API to get a random affirmation"
curl --silent --header "x-api-key: plZqPXiSRGA8b2mawMWNCfMPER9wVq1I" http://127.0.0.1/api/v2/random_affirmation | json_pp \
|| EXIT_STATUS=$?

echo ">>> Curling the API to get a list of affirmations"
curl --silent --header "x-api-key: plZqPXiSRGA8b2mawMWNCfMPER9wVq1I" http://127.0.0.1/api/v2/affirmations | json_pp \
|| EXIT_STATUS=$?

echo ">>> Curling the API to get a list of categories"
curl --silent --header "x-api-key: plZqPXiSRGA8b2mawMWNCfMPER9wVq1I" http://127.0.0.1/api/v2/categories | json_pp \
|| EXIT_STATUS=$?

if [ "$EXIT_STATUS" -eq "0" ]; then
echo ">>> OK"
else
echo ">>> ERROR"
fi

exit $EXIT_STATUS
