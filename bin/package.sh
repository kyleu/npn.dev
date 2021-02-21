#!/bin/bash

## Packages the build output for Github Releases

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..
pdir="$( pwd )"

# ./bin/build-all.sh

mkdir -p ./build/stage

rm -rf ./build/package
mkdir -p ./build/package

cd "$pdir/build/stage"

# cp -R "$pdir/data" ./data

pkg () {
  echo "$4 ($2)..."
  cp "$pdir/build/$1/$2/$3" "./$3"

  if [ $2 = "amd64" ]; then
    zip -r "$pdir/build/package/npn.server.$4.zip" *
  else
    zip -r "$pdir/build/package/npn.server.$4.$2.zip" *
  fi

  rm "./$3"
}

# macOS
pkg darwin amd64 npn macos
pkg darwin arm64 npn macos

echo "macOS app..."
cd ../darwin
zip -r "npn.app.macos.zip" npn.app
mv "npn.app.macos.zip" "../package"
cd ../stage

# Linux
pkg linux amd64 npn linux
pkg linux 386 npn linux
pkg linux arm64 npn linux
pkg linux arm npn linux
pkg linux mips npn linux
pkg linux riscv64 npn linux

# FreeBSD
pkg freebsd amd64 npn freebsd
pkg freebsd 386 npn freebsd
pkg freebsd arm64 npn freebsd
pkg freebsd arm npn freebsd

# Windows
pkg windows amd64 npn.exe windows
pkg windows 386 npn.exe windows
pkg windows arm npn.exe windows

# Docker
echo "docker..."
cp "$pdir/build/docker/npn.docker.tar.gz" "$pdir/build/package/npn.server.docker.tar.gz"

# WASM
echo "wasm..."
cp "$pdir/build/js/wasm/npn.wasm" ./npn.wasm
zip -r "$pdir/build/package/npn.server.wasm.zip" *
rm ./npn.wasm

# HTML
echo "html..."
pwd
cd "$pdir/projects/wasm/assets"
zip -r "$pdir/build/package/npn.server.html.zip" *
cd "$pdir/build/stage"

# Android
echo "android library..."
cp "$pdir/build/android/npn.aar" ./npn.aar
zip -r "$pdir/build/package/npn.library.android.zip" npn.aar
rm ./npn.aar

echo "android app..."
cp "$pdir/build/android/npn.apk" ./npn.apk
zip -r "$pdir/build/package/npn.app.android.zip" npn.apk
rm ./npn.apk

# iOS
echo "ios framework..."
cp  -r "$pdir/build/ios/NpnServer.framework" ./NpnServer.framework

cd NpnServer.framework
rm -rf Headers
rm -rf Modules
rm -rf NpnServer
rm -rf Resources
cp -R Versions/A/* .
rm -rf Versions
cd ..

zip -r "$pdir/build/package/npn.library.ios.zip" *
rm  -rf ./NpnServer.framework

echo "ios app..."
cd ../ios
zip -r "npn.app.ios.zip" npn.app
mv "npn.app.ios.zip" "../package"
cd ../stage

rm -rf "$pdir/build/stage"
