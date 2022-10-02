#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f "$0")")

rm -f $BASE_DIR/../bin/*
