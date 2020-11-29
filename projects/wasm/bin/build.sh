#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
dir="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$dir"

mkdir -p ./build

mkdir -p ./assets/vendor/uikit

cp ../../web/assets/vendor/npn.css ./assets/vendor/npn.css
cp ../../web/assets/vendor/npn.js ./assets/vendor/npn.js
cp ../../web/assets/vendor/vendor.js ./assets/vendor/vendor.js
cp ../../web/assets/vendor/editor/editor.js ./assets/vendor/editor/editor.js

cd cmd/wasm-npn
GOOS=js GOARCH=wasm go build -o ../../assets/npn.wasm

cd ../..
go build -o ./build/ ./cmd/wasm-test
