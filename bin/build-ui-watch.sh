#!/bin/bash

## Builds the TypeScript project, watching for changes
## Requires yarn available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../ui

../bin/build-ui.sh
echo "Watching TypeScript compilation for [ui]..."
yarn build
