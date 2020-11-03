#!/bin/bash

## Builds the project in release mode and runs it

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build-ui.sh
make build-release
build/release/npn
