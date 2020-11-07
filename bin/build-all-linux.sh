#!/bin/bash

## Builds all the linux-like apps

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build.sh linux amd64
# bin/build-webview.sh linux amd64

bin/build.sh linux 386
bin/build.sh linux arm64
bin/build.sh linux arm
bin/build.sh linux mips
bin/build.sh linux riscv64

bin/build.sh freebsd amd64
bin/build.sh freebsd 386
bin/build.sh freebsd arm64
bin/build.sh freebsd arm
