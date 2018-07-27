#!/usr/bin/env bash

set -o nounset
set -o errexit

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${DIR}

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -a -o ~/bin/bullettrain
