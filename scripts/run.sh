#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f "$0")")
CONTAINER_NAME=rafapi
DATA_SOURCE=db/db.dist.sqlite

docker rm -f $CONTAINER_NAME &> /dev/null && echo ">>> previous instance of $CONTAINER_NAME was stopped and removed"
docker run \
-d \
--name $CONTAINER_NAME \
-p 5000:5000 \
--restart unless-stopped \
-v $BASE_DIR/../$DATA_SOURCE:/app/$DATA_SOURCE \
$CONTAINER_NAME
