#!/bin/bash

## Builds the project as an iOS framework and builds the native app in `projects/ios`

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "Building [ios]..."

mkdir -p build/ios/

bin/asset-embed.sh
echo "gomobile..."
gomobile bind -o build/ios/NpnServer.framework -target=ios github.com/kyleu/npn/lib
bin/asset-reset.sh

echo "Building [ios] app..."

cd projects/ios/npn

xcodebuild -project npn.xcodeproj

cd build/Release-iphoneos/

cp -R npn.app ../../../../../build/ios
