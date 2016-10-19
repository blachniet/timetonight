#!/bin/bash
export GOOS=linux
export GOARCH=amd64
go build ../cmd/timetonight/...
cp -r ../templates .
docker build -t blachniet/timetonight .