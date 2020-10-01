#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

mkdir -p ./build

mkdir -p ./assets/vendor/uikit

cp ../../web/assets/vendor/npn.css ./assets/vendor/npn.css
cp ../../web/assets/vendor/npn.js ./assets/vendor/npn.js

cp ../../npnasset/vendor/uikit/uikit.min.css ./assets/vendor/uikit/uikit.min.css
cp ../../npnasset/vendor/uikit/uikit.min.js ./assets/vendor/uikit/uikit.min.js
cp ../../npnasset/vendor/uikit/uikit-icons.min.js ./assets/vendor/uikit/uikit-icons.min.js

cd cmd/wasm-npn
GOOS=js GOARCH=wasm go build -o ../../assets/npn.wasm

cd ../..
go build -o ./build/ ./cmd/wasm-test
