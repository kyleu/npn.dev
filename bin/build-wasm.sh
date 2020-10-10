#!/bin/bash

## Builds the app as a WASM server

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build.sh js wasm
mv ./build/js/wasm/npn ./build/js/wasm/npn.wasm
