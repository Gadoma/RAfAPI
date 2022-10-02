#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f "$0")")

cd $BASE_DIR/../
go test -count=1 ./...
