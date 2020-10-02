#!/bin/bash

## Uses `tsc` to compile the scripts in `client`
## Requires tsc available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../client

echo "Building TypeScript..."
tsc --project tsconfig.json

cd $project_dir/web/assets/vendor

echo "Optimizing TypeScript..."
closure-compiler --create_source_map npn.min.js.map npn.js > npn.min.js
