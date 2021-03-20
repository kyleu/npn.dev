#!/bin/bash

## Builds the screenshots from various platforms

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

cd projects/screenshots
echo "Building screnshots..."
node index.js
