#!/bin/bash

## Attempts to build for all available platforms and architectures; Requires docker and a bunch of other stuff

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build-desktop.sh

bin/build-docker.sh

bin/build-wasm.sh
projects/wasm/bin/build.sh

bin/build-android.sh
bin/build-ios.sh

