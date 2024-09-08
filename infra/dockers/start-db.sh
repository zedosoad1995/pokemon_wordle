#!/bin/sh
set -e

docker-compose --file="infra/dockers/docker-compose-start-db.yml" -p pokemon_wordle up -d