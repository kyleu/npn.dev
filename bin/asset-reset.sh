#!/bin/bash

## Resets the assets to load from local filesystem in development

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "Resetting assets..."
git checkout app/assets/assets.go
