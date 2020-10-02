#!/bin/bash

## Builds the project as a macOS server and builds the native app in `projects/macos`

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "Building [macos] app..."

mkdir -p build/darwin/

bin/build.sh darwin amd64
cp build/darwin/amd64/npn projects/macos/npn/npn/npn-server

cd projects/macos/npn

xcodebuild -project npn.xcodeproj

cd build/Release/

cp -R npn.app ../../../../../build/darwin
