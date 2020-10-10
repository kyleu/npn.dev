#!/bin/bash

## Builds the project as an android framework and builds the native app in `projects/android`

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

pdir="$( pwd )"

echo "Building [android]..."

mkdir -p build/android/

bin/asset-embed.sh
echo "gomobile..."
gomobile bind -o build/android/npn.aar -target=android github.com/kyleu/npn/lib
bin/asset-reset.sh

cd projects/android/npn/app/libs
rm -f npn.aar npn-sources.jar
cp ${pdir}/build/android/npn.aar .
cp ${pdir}/build/android/npn-sources.jar .

cd "${pdir}/projects/android/npn"

gradle assembleDebug

cp "app/build/outputs/apk/debug/app-debug.apk" "${pdir}/build/android/npn.apk"

cd "${pdir}/build/android"
