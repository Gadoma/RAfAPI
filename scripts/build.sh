#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f "$0")")
CONTAINER_NAME=rafapi

docker build -t $CONTAINER_NAME -f $BASE_DIR/../docker/Dockerfile $BASE_DIR/../
