#!/bin/bash

## Attempts to build the wasm app

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

mkdir -p ./build
mkdir -p ./assets/vendor/editor

cp ../../web/assets/vendor/npn.css ./assets/vendor/npn.css
cp ../../web/assets/vendor/npn.js ./assets/vendor/npn.js
cp ../../web/assets/vendor/vendor.js ./assets/vendor/vendor.js
cp ../../web/assets/vendor/editor/editor.js ./assets/vendor/editor/editor.js

pwd
cd cmd/wasm-npn
GOOS=js GOARCH=wasm go build -o ../../assets/npn.wasm

cd ../..
go build -o ./build/ ./cmd/wasm-test
