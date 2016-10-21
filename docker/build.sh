#!/bin/bash
export GOOS=linux
export GOARCH=amd64
go build ../cmd/timetonight/...
cp -r ../templates .
cp /usr/local/go/lib/time/zoneinfo.zip .
docker build -t blachniet/timetonight .
