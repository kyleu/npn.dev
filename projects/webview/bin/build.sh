#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
dir="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$dir"

os=${1:-darwin}
arch=${2:-amd64}
fn=${3:-npn-webview}

echo "Building webview for [$os $arch]..."
env GOOS=$os GOARCH=$arch make build-release

mkdir -p ./build/$os/$arch
mv "./build/release/$fn" "./build/$os/$arch/$fn"
