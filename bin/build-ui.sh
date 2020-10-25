#!/bin/bash

## Uses `./ui/bin/copy.sh` to build the project and copy files to the main app

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../ui

echo "Building UI..."
bin/copy.sh
