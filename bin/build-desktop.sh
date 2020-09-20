#!/bin/bash

## Builds the desktop apps

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

bin/build-macos.sh

# bin/build.sh darwin arm64

bin/build.sh linux amd64
bin/build.sh linux 386
bin/build.sh linux arm64
bin/build.sh linux arm
bin/build.sh linux mips
bin/build.sh linux riscv64

bin/build.sh freebsd amd64
bin/build.sh freebsd 386
bin/build.sh freebsd arm64

bin/build.sh windows amd64 npn.exe
bin/build.sh windows 386 npn.exe
bin/build.sh windows arm npn.exe
# bin/build.sh windows arm64 npn.exe
