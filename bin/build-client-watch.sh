#!/bin/bash

## Builds the TypeScript resources using `build-client`, then watches for changes
## Requires tsc available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir/client

../bin/build-client.sh
echo "Watching TypeScript compilation for [client/src]..."
tsc -w --project tsconfig.json
