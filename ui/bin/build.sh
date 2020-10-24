#!/bin/bash

## Builds the TypeScript project
## Requires yarn available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

yarn build
