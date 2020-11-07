#!/bin/bash

## Builds all the apps that can only be built on macOS

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build-macos.sh
bin/build-ios.sh
