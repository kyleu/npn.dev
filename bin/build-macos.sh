#!/bin/bash

## Builds the project as a macOS server and builds the native app in `projects/macos`

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

mkdir -p build/darwin/

bin/build.sh darwin arm64
bin/build.sh darwin amd64

echo "Building [macos] app..."
cp build/darwin/amd64/npn projects/macos/npn/npn/npn-server
cd projects/macos/npn
xcodebuild -project npn.xcodeproj
cd build/Release/
cp -R npn.app ../../../../../build/darwin
