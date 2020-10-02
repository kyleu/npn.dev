#!/bin/bash

## Resets the assets to load from local filesystem in development

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "Resetting assets..."
git checkout npnasset/assets/assets.go
git checkout app/assets/assets.go
