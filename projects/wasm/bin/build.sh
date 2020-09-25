#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

mkdir -p ./build

go build -o ./build/ ./cmd/wasm-test
GOOS=js GOARCH=wasm go build -o build/npn.wasm github.com/kyleu/npn/projects/wasm/cmd/wasm-npn

cp build/npn.wasm assets/
