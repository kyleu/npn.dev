#!/bin/bash

## Builds all the desktop apps, XCode required

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build-macos.sh
# bin/build-webview.sh darwin amd64

# bin/build.sh darwin arm64

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

bin/build.sh windows amd64 npn.exe
# bin/build-webview.sh windows amd64

bin/build.sh windows 386 npn.exe
bin/build.sh windows arm npn.exe
# bin/build.sh windows arm64 npn.exe
