#!/bin/bash

## Uses `tsc` to compile the scripts in `client`
## Requires tsc available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir/client

tsc --project tsconfig.json

cd $project_dir/web/assets

closure-compiler --create_source_map npn.min.js.map npn.js > npn.min.js
