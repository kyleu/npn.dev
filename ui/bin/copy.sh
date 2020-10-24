#!/bin/bash

## Copies build files to main project

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build.sh

mkdir -p ../web/assets/vendor/ui

rm -rf ../web/assets/vendor/ui/*

cp dist/js/app.*.js ../web/assets/vendor/ui/npn.js
cp dist/js/chunk*.*.js ../web/assets/vendor/ui/vendor.js
cp dist/js/*.map ../web/assets/vendor/ui
cp dist/css/app.*.css ../web/assets/vendor/ui/npn.css
