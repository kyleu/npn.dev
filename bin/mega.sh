#!/bin/bash

## Mega!

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

rm -rf build
bin/build-all.sh
bin/package.sh
bin/deploy.sh
