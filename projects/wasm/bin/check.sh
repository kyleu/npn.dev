#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "vetting [wasm-test]"
cd cmd/wasm-test
go vet

echo "vetting [wasm-npn]"
cd ../../
cd cmd/wasm-npn
go vet
