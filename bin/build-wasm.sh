#!/bin/bash

## Builds the desktop apps

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

bin/build.sh js wasm
mv ./build/js/wasm/npn ./build/js/wasm/npn.wasm
