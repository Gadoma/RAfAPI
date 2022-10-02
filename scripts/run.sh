#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f "$0")")
CONTAINER_NAME=rafapi

docker rm -f $CONTAINER_NAME &> /dev/null && echo ">>> previous instance of $CONTAINER_NAME was stopped and removed"
docker run -d -p 5000:5000 --name $CONTAINER_NAME --restart unless-stopped $CONTAINER_NAME
