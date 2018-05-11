#!/usr/bin/env bash

rm -rfv ./releases/*

# go tool dist list

# Raspberry pi 1?
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "-s -w" -a -o ./releases/bullettrain.linux-armv6
sha256sum -b ./releases/bullettrain.linux-armv6 > ./releases/bullettrain.linux-armv6.sha256

# Raspberry pi 2?
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-s -w" -a -o ./releases/bullettrain.linux-armv7
sha256sum -b ./releases/bullettrain.linux-armv7 > ./releases/bullettrain.linux-armv7.sha256

# Raspberry pi 3
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -a -o ./releases/bullettrain.linux-arm64
sha256sum -b ./releases/bullettrain.linux-arm64 > ./releases/bullettrain.linux-arm64.sha256

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -a -o ./releases/bullettrain.linux-amd64
sha256sum -b ./releases/bullettrain.linux-amd64 > ./releases/bullettrain.linux-amd64.sha256

# Mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -a -o ./releases/bullettrain.darwin-amd64
sha256sum -b ./releases/bullettrain.darwin-amd64 > ./releases/bullettrain.darwin-amd64.sha256
