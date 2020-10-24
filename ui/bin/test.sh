#!/bin/bash

## Builds the TypeScript project, watching for changes
## Requires yarn available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

yarn build

cd dist
http.py 10102
