#!/bin/sh
set -e

docker-compose --file="infra/dockers/docker-compose-start-db.yml" -p pokegrid up -d