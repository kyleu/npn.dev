#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

mkdir -p ./build

go build -o ./build/ ./cmd/wasm-test

cd cmd/wasm-npn
GOOS=js GOARCH=wasm go build -o ../../assets/npn.wasm
