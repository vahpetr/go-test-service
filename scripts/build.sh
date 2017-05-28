#!/bin/bash

# chmod -R +x ./scripts

set -e

docker rm -f go-test-service || true
docker rmi -f go-test-service || true

cd src

export GOOS=linux GOARCH=amd64 CGO_ENABLED=0

go build -a -installsuffix cgo -o service .

cd -

docker build -t go-test-service -f Dockerfile .

docker images | grep go-test-service

# docker run -d -p 8080:8080 --name go-test-service go-test-service
