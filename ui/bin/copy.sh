#!/bin/bash

## Copies build files to main project

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build.sh

mkdir -p ../web/assets/vendor

rm -f ../web/assets/vendor/*.js
rm -f ../web/assets/vendor/*.css
rm -f ../web/assets/vendor/*.map

cp dist/js/app.*.js ../web/assets/vendor/npn.js
cp dist/js/chunk*.*.js ../web/assets/vendor/vendor.js
cp dist/js/*.map ../web/assets/vendor
cp dist/css/app.*.css ../web/assets/vendor/npn.css
