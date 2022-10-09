#!/usr/bin/env bash

docker compose -f docker-compose.dist.yml --env-file .env.dist --profile app build
