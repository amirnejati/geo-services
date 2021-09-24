#!/bin/bash

export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin

# SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"
cd "$SCRIPT_DIR" || return

PROTO_FILE=${1:-"geofence.proto"}
protoc -I. "$PROTO_FILE" --go_out=plugins=grpc:.
